package dump

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DumpCommand = &cobra.Command{
	Use:   "dump",
	Short: "Dump all sites and entries in the specificed format",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
	},
}
