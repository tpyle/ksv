package sites

import "github.com/spf13/cobra"

func init() {
	SitesCommand.AddCommand(addSiteCmd)
}

var SitesCommand = &cobra.Command{
	Use:   "sites",
	Short: "Manage sites",
}
