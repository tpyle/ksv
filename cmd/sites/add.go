package sites

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/util"
	"github.com/tpyle/ksv/lib/types"
)

var (
	siteUrl       string
	siteName      string
	siteAppId     string
	siteGenericId string
	siteNotes     string
)

func init() {
	SitesCommand.AddCommand(addSiteCmd)

	addSiteCmd.Flags().StringVarP(&siteUrl, "url", "u", "", "Site URL")
	addSiteCmd.Flags().StringVarP(&siteName, "name", "s", "", "Site Name")
	addSiteCmd.Flags().StringVarP(&siteAppId, "appId", "a", "", "Site App ID")
	addSiteCmd.Flags().StringVarP(&siteGenericId, "genericId", "g", "", "Site Generic ID")
	addSiteCmd.Flags().StringVarP(&siteNotes, "notes", "o", "", "Site Notes")
	addSiteCmd.MarkFlagRequired("name")
	addSiteCmd.MarkFlagsOneRequired("url", "appId", "genericId")
}

var addSiteCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new site",
	Run: func(cmd *cobra.Command, args []string) {
		site := types.NewSite(siteUrl, siteName, siteAppId, siteGenericId, siteNotes)
		if errs := site.Validate(); len(errs) > 0 {
			logrus.Errorf("Validation failed")
			for _, err := range errs {
				logrus.Errorf("  %s", err)
			}
			return
		}

		ksv, err := util.GetKeystoreFromContext(cmd.Context())
		if err != nil {
			logrus.WithError(err).Fatal("Failed to get keystore")
		}

		err = ksv.DefaultNamespace.AddSite(site)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to add site")
		}

		err = util.SaveKeystoreFromContext(cmd.Context(), ksv)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to save keystore")
		}
	},
}
