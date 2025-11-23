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
	addSiteCmd.Flags().StringVar(&siteUrl, "url", "", "Site URL")
	addSiteCmd.Flags().StringVar(&siteName, "name", "", "Site Name")
	addSiteCmd.Flags().StringVar(&siteAppId, "appId", "", "Site App ID")
	addSiteCmd.Flags().StringVar(&siteGenericId, "genericId", "", "Site Generic ID")
	addSiteCmd.Flags().StringVar(&siteNotes, "notes", "", "Site Notes")

	SitesCommand.AddCommand(addSiteCmd)
}

var addSiteCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new site",
	Run: func(cmd *cobra.Command, args []string) {
		site := types.Site{
			Url:       types.NewKSVString(siteUrl),
			Name:      types.NewKSVString(siteName),
			AppId:     types.NewKSVString(siteAppId),
			GenericId: types.NewKSVString(siteGenericId),
			Notes:     types.NewKSVString(siteNotes),
		}
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

		if ksv.DefaultNamespace.Sites == nil {
			ksv.DefaultNamespace.Sites = make(types.KSVMap[*types.Site])
		}
		if _, ok := ksv.DefaultNamespace.Sites[siteName]; ok {
			logrus.Fatalf("Site with name %s already exists", siteName)
		}

		ksv.DefaultNamespace.Sites[siteName] = &site

		err = util.SaveKeystoreFromContext(cmd.Context(), ksv)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to save keystore")
		}
	},
}
