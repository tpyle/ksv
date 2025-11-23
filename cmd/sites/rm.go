package sites

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/util"
)

func init() {
	rmSiteCmd.Flags().StringVar(&siteName, "name", "", "Site Name")

	SitesCommand.AddCommand(rmSiteCmd)
}

var rmSiteCmd = &cobra.Command{
	Use:   "rm [site]",
	Short: "Remove a site",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		ksv, err := util.GetKeystoreFromContext(cmd.Context())
		if err != nil {
			logrus.WithError(err).Fatal("Failed to get keystore")
		}

		if ksv.DefaultNamespace.Sites == nil {
			logrus.Fatal("No sites found in the keystore")
		}

		siteName := args[0]

		if _, ok := ksv.DefaultNamespace.Sites[siteName]; ok {
			// TODO: Prompt for correctness

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
