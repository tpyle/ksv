package entries

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/util"
)

func init() {
	EntriesCommand.AddCommand(rmEntryCmd)
	rmEntryCmd.Flags().StringVarP(&siteName, "site-name", "s", "", "Site Name")
	rmEntryCmd.Flags().StringVarP(&entryName, "entry-name", "e", "", "Entry Name")
	rmEntryCmd.MarkFlagRequired("site-name")
	rmEntryCmd.MarkFlagRequired("entry-name")
}

var rmEntryCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove an entry from a site",
	Run: func(cmd *cobra.Command, args []string) {
		ksv, err := util.GetKeystoreFromContext(cmd.Context())
		if err != nil {
			logrus.WithError(err).Fatal("Failed to get keystore")
		}

		site, ok := ksv.DefaultNamespace.Sites[siteName]
		if !ok {
			logrus.Fatalf("Site with name %s does not exist", siteName)
		}

		err = site.RemoveEntry(entryName)
		if err != nil {
			logrus.WithError(err).Fatalf("Failed to remove entry %s from site %s", entryName, siteName)
		}

		err = util.SaveKeystoreFromContext(cmd.Context(), ksv)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to save keystore")
		}

		logrus.Infof("Entry %s removed from site %s", entryName, siteName)
	},
}
