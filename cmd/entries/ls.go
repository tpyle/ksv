package entries

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/util"
)

func init() {
	EntriesCommand.AddCommand(lsEntryCmd)
	lsEntryCmd.Flags().StringVarP(&siteName, "site-name", "s", "", "Site Name")
	lsEntryCmd.MarkFlagRequired("site-name")
}

var lsEntryCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List entries for a site",
	Run: func(cmd *cobra.Command, args []string) {
		ksv, err := util.GetKeystoreFromContext(cmd.Context())
		if err != nil {
			logrus.WithError(err).Fatal("Failed to get keystore")
		}

		site, ok := ksv.DefaultNamespace.Sites[siteName]
		if !ok {
			logrus.Fatalf("Site with name %s does not exist", siteName)
		}

		for entryName := range site.Entries {
			fmt.Printf("%s\n", entryName)
		}
	},
}
