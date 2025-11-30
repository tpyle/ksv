package search

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/util"
	"github.com/tpyle/ksv/lib/types"
)

var (
	useRegex bool
	ksv      *types.KSV
)

func init() {
	SearchCommand.Flags().BoolVarP(&useRegex, "regex", "r", false, "Use regex for searching")
}

func maxlen(arr []string) int {
	max := 0
	for _, s := range arr {
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}

var SearchCommand = &cobra.Command{
	Use:   "search [query]",
	Short: "Search through your sites and entries",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Pre-run logic can be added here if needed
		var err error
		ksv, err = util.GetKeystoreFromContext(cmd.Context())
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		var match *regexp.Regexp
		if useRegex {
			var err error
			match, err = regexp.Compile(query)
			if err != nil {
				logrus.WithError(err).Fatal("Invalid regex")
				return
			}
		}
		values := ksv.DefaultNamespace.GetValues()
		keys := []string{}
		for k := range values {
			if !useRegex {
				if strings.Contains(k, args[0]) {
					keys = append(keys, k)
				}
			} else {
				if match.MatchString(k) {
					keys = append(keys, k)
				}
			}
		}
		sort.Strings(keys)
		columnWidth := maxlen(keys) + 2
		for _, k := range keys {
			fmt.Printf("%-*s \t%v\n", columnWidth, k, values[k])
		}
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		// No need to save keystore after search as it's read-only
		return nil
	},
}
