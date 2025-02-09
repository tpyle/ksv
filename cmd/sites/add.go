package sites

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/lib/types"
)

var (
	siteUrl   string
	siteName  string
	siteAppId string
	siteNotes string
)

func init() {
	addSiteCmd.Flags().StringVar(&siteUrl, "url", "", "Site URL")
	addSiteCmd.Flags().StringVar(&siteName, "name", "", "Site Name")
	addSiteCmd.Flags().StringVar(&siteAppId, "appId", "", "Site App ID")
	addSiteCmd.Flags().StringVar(&siteNotes, "notes", "", "Site Notes")
}

var addSiteCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new site",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to add a new site
		site := types.Site{
			Url:   siteUrl,
			Name:  siteName,
			AppId: siteAppId,
			Notes: siteNotes,
		}
		// Add site to the KSV structure (this is a placeholder, actual implementation may vary)
		fmt.Printf("Site added: %+v\n", site)
	},
}
