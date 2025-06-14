package sites

import "github.com/spf13/cobra"

func init() {
	SitesCommand.AddCommand(addSiteCmd)
	SitesCommand.AddCommand(rmSiteCmd)
}

var SitesCommand = &cobra.Command{
	Use:   "sites",
	Short: "Manage sites",
}
