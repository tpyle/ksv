package entries

import (
	"github.com/spf13/cobra"
)

var (
	siteUrl   string
	siteName  string
	siteAppId string
	siteNotes string

	entryName     string
	entryUsername string
	entryEmail    string
	entryNotes    string
)

func init() {
	addEntryCmd.Flags().StringVar(&entryName, "name", "", "Entry Name")
	addEntryCmd.Flags().StringVar(&entryUsername, "username", "", "Entry Username")
	addEntryCmd.Flags().StringVar(&entryEmail, "email", "", "Entry Email")
	addEntryCmd.Flags().StringVar(&entryNotes, "notes", "", "Entry Notes")
}

var addEntryCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
