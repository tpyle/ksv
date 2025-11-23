package sites

import "github.com/spf13/cobra"

var SitesCommand = &cobra.Command{
	Use: "sites",
	Aliases: []string{
		"site",
	},
	Short: "Manage sites",
}
