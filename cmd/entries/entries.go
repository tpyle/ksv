package entries

import "github.com/spf13/cobra"

func init() {
	EntriesCommand.AddCommand(addEntryCmd)
}

var EntriesCommand = &cobra.Command{
	Use:   "entries",
	Short: "Manage entries",
}
