package entries

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/util"
	"github.com/tpyle/ksv/lib/types"
	"golang.org/x/term"
)

var (
	siteName           string
	entryName          string
	entryUsername      string
	entryEmail         string
	entryNotes         string
	entryIdp           string
	customFields       []string
	secretFields       []string
	promptSecretFields bool
)

func init() {
	EntriesCommand.AddCommand(addEntryCmd)

	addEntryCmd.Flags().StringVarP(&siteName, "site-name", "s", "", "Site Name")
	addEntryCmd.Flags().StringVarP(&entryName, "entry-name", "e", "", "Entry Name")
	addEntryCmd.Flags().StringVarP(&entryUsername, "username", "u", "", "Entry Username")
	addEntryCmd.Flags().StringVarP(&entryEmail, "email", "m", "", "Entry Email")
	addEntryCmd.Flags().StringVarP(&entryNotes, "notes", "o", "", "Entry Notes")
	addEntryCmd.Flags().StringVar(&entryIdp, "idp", "", "Identity Provider (IdP) name")
	addEntryCmd.Flags().StringArrayVar(&customFields, "custom", []string{}, "Custom fields in the format key=value")
	addEntryCmd.Flags().StringArrayVar(&secretFields, "secret", []string{}, "Secret fields in the format key=value")
	addEntryCmd.Flags().BoolVarP(&promptSecretFields, "prompt-secret", "p", false, "Prompt for secret fields interactively")
	addEntryCmd.MarkFlagsMutuallyExclusive("secret", "prompt-secret")
	addEntryCmd.MarkFlagRequired("site-name")
	addEntryCmd.MarkFlagRequired("entry-name")
	addEntryCmd.MarkFlagsOneRequired("username", "email", "custom", "secret", "prompt-secret")
}

func ParseField(input string) (string, string, error) {
	parts := strings.SplitN(input, "=", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid field format: %s: should be <key>=<value>", input)
	}
	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	if key == "" {
		return "", "", fmt.Errorf("invalid field format: %s: key cannot be empty", input)
	}
	return key, value, nil
}

func ReadInteractiveSecrets() map[string]string {
	secrets := make(map[string]string)
	fmt.Println("Please enter each secret field, in the format <key>=<value>. Press Ctrl+D (or Ctrl+Z on Windows) or enter an empty line to finish.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if term.IsTerminal(int(os.Stdin.Fd())) {
			fmt.Print("> ")
		}
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if input == "" {
			break
		}
		key, value, err := ParseField(input)
		if err != nil {
			logrus.Error(err)
			continue
		}
		secrets[key] = value
	}
	return secrets
}

var addEntryCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry",
	Run: func(cmd *cobra.Command, args []string) {
		ksv, err := util.GetKeystoreFromContext(cmd.Context())
		if err != nil {
			logrus.WithError(err).Fatal("Failed to get keystore")
		}

		site, ok := ksv.DefaultNamespace.Sites[siteName]
		if !ok {
			logrus.Fatalf("Site with name %s does not exist", siteName)
		}

		if entryIdp != "" {
			_, ok = ksv.DefaultNamespace.GetIDPByName(entryIdp)
			if !ok {
				logrus.Fatalf("IDP with name %s does not exist", entryIdp)
			}
		}

		entry := types.NewEntry(entryUsername, entryEmail, entryNotes, entryIdp != "", entryIdp)

		if site.HasEntry(entryName) {
			logrus.Fatalf("Entry with name %s already exists in site %s", entryName, siteName)
		}

		if promptSecretFields {
			secrets := ReadInteractiveSecrets()
			for k, v := range secrets {
				entry.SecretFields[k] = types.NewKSVString(v)
			}
		} else {
			for _, field := range secretFields {
				key, value, err := ParseField(field)
				if err != nil {
					logrus.WithError(err).Fatal("Failed to parse secret field")
				}
				entry.SecretFields[key] = types.NewKSVString(value)
			}
		}

		for _, field := range customFields {
			key, value, err := ParseField(field)
			if err != nil {
				logrus.WithError(err).Fatal("Failed to parse custom field")
			}
			entry.CustomFields[key] = types.NewKSVString(value)
		}

		err = site.AddEntry(entryName, entry)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to add entry to site")
		}

		err = util.SaveKeystoreFromContext(cmd.Context(), ksv)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to save keystore")
		}
	},
}
