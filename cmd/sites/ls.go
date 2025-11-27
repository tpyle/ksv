package sites

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/util"
	"github.com/tpyle/ksv/lib/ksverrors"
)

func init() {
	SitesCommand.AddCommand(lsSiteCmd)
}

var lsSiteCmd = &cobra.Command{
	Use: "list",
	Aliases: []string{
		"ls",
	},
	Short: "List all sites",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return ksverrors.ErrCliNoMoreArgs
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ksv, err := util.GetKeystoreFromContext(cmd.Context())
		if err != nil {
			return ksverrors.ErrCliFailedToGetKeystore
		}

		if ksv.DefaultNamespace.Sites != nil {
			for _, site := range ksv.DefaultNamespace.Sites {
				fmt.Println(site.Name.Value)
			}
		}
		return nil
	},
}
