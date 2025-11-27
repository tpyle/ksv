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

	addEntryCmd.Flags().StringVarP(&siteName, "site-name", "s", "", "Site Name")
	addEntryCmd.Flags().StringVarP(&entryName, "entry-name", "e", "", "Entry Name")
	addEntryCmd.Flags().StringVarP(&entryUsername, "username", "u", "", "Entry Username")
	addEntryCmd.Flags().StringVarP(&entryEmail, "email", "m", "", "Entry Email")
	addEntryCmd.Flags().StringVarP(&entryNotes, "notes", "o", "", "Entry Notes")
	addEntryCmd.Flags().StringVar(&entryIdp, "idp", "", "Identity Provider (IdP) name")
	addEntryCmd.Flags().StringArrayVar(&customFields, "custom", []string{}, "Custom fields in the format key=value")
	addEntryCmd.Flags().StringArrayVar(&secretFields, "secret", []string{}, "Secret fields in the format key=value")
	addEntryCmd.Flags().BoolVar(&promptSecretFields, "prompt-secret", false, "Prompt for secret fields interactively")
	addEntryCmd.MarkFlagsMutuallyExclusive("secret", "prompt-secret")
	addEntryCmd.MarkFlagRequired("site-name")
	addEntryCmd.MarkFlagRequired("entry-name")
	addEntryCmd.MarkFlagsOneRequired("username", "email", "custom", "secret", "prompt-secret")
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
