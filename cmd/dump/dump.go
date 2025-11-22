package dump

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/util"
	"github.com/tpyle/ksv/lib/ksverrors"
	"sigs.k8s.io/yaml"
)

const (
	DumpPadding = 3 // Padding for table output
)

var (
	dumpFormat  string
	skipHeaders bool
)

func init() {
	DumpCommand.Flags().StringVarP(&dumpFormat, "output", "o", "table", "Output format (json, yaml, table)")
	DumpCommand.Flags().BoolVar(&skipHeaders, "no-headers", false, "Disable headers in table output")
}

var DumpCommand = &cobra.Command{
	Use:   "dump",
	Short: "Dump all data in the specificed format",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			logrus.Fatal("No additional arguments are allowed")
		}
		if dumpFormat != "json" && dumpFormat != "yaml" && dumpFormat != "table" {
			logrus.Fatalf("Unsupported format: %s. Supported formats are: json, yaml, table", dumpFormat)
		}
		if skipHeaders && dumpFormat != "table" {
			logrus.Fatal("Headers can only be used with table format")
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ksv, err := util.GetKeystoreFromContext(cmd.Context())
		if err != nil {
			logrus.WithError(err).Fatal("Failed to get keystore")
		}

		switch dumpFormat {
		case "table":
			// Print the keystore in a human-readable table format, with formatted columns
			// Columns are Site (Name), Entry (Username)
			if ksv.DefaultNamespace.Sites != nil {
				siteColumnWidth := 20
				entryColumnWidth := 30
				for _, site := range ksv.DefaultNamespace.Sites {
					if len(site.Name.Value) > siteColumnWidth {
						siteColumnWidth = len(site.Name.Value) + DumpPadding
					}
					if site.Entries != nil {
						for _, entry := range site.Entries {
							if len(entry.Username.Value) > entryColumnWidth {
								entryColumnWidth = len(entry.Username.Value) + DumpPadding
							}
						}
					}
				}

				// Print header (e.g. "Site Name        Entry Username         ")
				if !skipHeaders {
					fmt.Printf("%-*s %-*s\n", siteColumnWidth, "Site Name", entryColumnWidth, "Entry Username")
				}
				// Print each site and its entries
				for _, site := range ksv.DefaultNamespace.Sites {
					if site.Entries != nil {
						for _, entry := range site.Entries {
							fmt.Printf("%-*s ", siteColumnWidth, site.Name.Value)
							fmt.Printf("%-*s ", entryColumnWidth, entry.Username.Value)
							fmt.Println()
						}
					}
				}
			}
		case "json":
			err := json.NewEncoder(cmd.OutOrStdout()).Encode(ksv)
			if err != nil {
				return ksverrors.ErrCliCouldNotEncodeJSON
			}
		case "yaml":
			bytes, err := yaml.Marshal(ksv)
			if err != nil {
				return ksverrors.ErrCliCouldNotEncodeYaml
			}
			_, err = cmd.OutOrStdout().Write(bytes)
			if err != nil {
				return ksverrors.ErrCliCouldNotWrite
			}
		}
		return nil
	},
}
