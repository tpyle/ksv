package search

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SearchCommand = &cobra.Command{
	Use:   "search",
	Short: "Interactively search through your sites and entries",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
	},
}
