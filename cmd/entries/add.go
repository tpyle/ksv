package entries

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	siteName           string
	entryName          string
	entryUsername      string
	entryEmail         string
	entryNotes         string
	customFields       []string
	secretFields       []string
	promptSecretFields bool
)

func init() {
	EntriesCommand.AddCommand(addEntryCmd)

	addEntryCmd.Flags().StringVarP(&siteName, "site", "s", "", "Site Name")
	addEntryCmd.Flags().StringVarP(&entryName, "name", "n", "", "Entry Name")
	addEntryCmd.Flags().StringVar(&entryUsername, "username", "", "Entry Username")
	addEntryCmd.Flags().StringVar(&entryEmail, "email", "", "Entry Email")
	addEntryCmd.Flags().StringVar(&entryNotes, "notes", "", "Entry Notes")
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
	},
}
