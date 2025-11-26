package entries

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/util"
	"github.com/tpyle/ksv/lib/types"
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

	addEntryCmd.Flags().StringVarP(&siteName, "site", "s", "", "Site Name")
	addEntryCmd.Flags().StringVar(&entryName, "name", "", "Entry Name")
	addEntryCmd.Flags().StringVar(&entryUsername, "username", "", "Entry Username")
	addEntryCmd.Flags().StringVar(&entryEmail, "email", "", "Entry Email")
	addEntryCmd.Flags().StringVar(&entryNotes, "notes", "", "Entry Notes")
	addEntryCmd.Flags().StringVar(&entryIdp, "idp", "", "Identity Provider (IdP) name")
	addEntryCmd.Flags().StringArrayVar(&customFields, "custom", []string{}, "Custom fields in the format key=value")
	addEntryCmd.Flags().StringArrayVar(&secretFields, "secret", []string{}, "Secret fields in the format key=value")
	addEntryCmd.Flags().BoolVar(&promptSecretFields, "prompt-secret", false, "Prompt for secret fields interactively")
}

var addEntryCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry",
	PreRun: func(cmd *cobra.Command, args []string) {
		if siteName == "" {
			logrus.Fatal("Site name is required")
		}
		if entryName == "" {
			logrus.Fatal("Entry name is required")
		}
		if entryUsername == "" && entryEmail == "" && entryNotes == "" && len(customFields) == 0 && len(secretFields) == 0 {
			logrus.Fatal("At least one field (username, email, notes, custom fields, or secret fields) must be provided")
		}
		if promptSecretFields && len(secretFields) > 0 {
			logrus.Fatal("Cannot use --prompt-secret with --secret fields")
		}
	},
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
