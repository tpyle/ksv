package sites

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/util"
)

func init() {
	rmSiteCmd.Flags().StringVar(&siteName, "name", "", "Site Name")
}

var rmSiteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a site",
	PreRun: func(cmd *cobra.Command, args []string) {
		if siteName == "" {
			logrus.Fatal("Site name must be provided")
		}
		if len(args) > 0 {
			logrus.Fatal("No additional arguments are allowed")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		ksv, err := util.GetKeystoreFromContext(cmd.Context())
		if err != nil {
			logrus.WithError(err).Fatal("Failed to get keystore")
		}

		if ksv.DefaultNamespace.Sites == nil {
			logrus.Fatal("No sites found in the keystore")
		}
		if _, ok := ksv.DefaultNamespace.Sites[siteName]; ok {
			delete(ksv.DefaultNamespace.Sites, siteName)
		} else {
			logrus.Fatalf("Site with name %s does not exist", siteName)
		}

		err = util.SaveKeystoreFromContext(cmd.Context(), ksv)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to save keystore")
		}
	},
}
