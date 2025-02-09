package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/config"
	"github.com/tpyle/ksv/cmd/dump"
	"github.com/tpyle/ksv/cmd/entries"
	"github.com/tpyle/ksv/cmd/search"
	"github.com/tpyle/ksv/cmd/sites"
)

var rootCmd = &cobra.Command{
	Use:   "ksv",
	Short: "KSV is a CLI tool for managing secrets",
}

func init() {
	rootCmd.AddCommand(sites.SitesCommand)
	rootCmd.AddCommand(search.SearchCommand)
	rootCmd.AddCommand(dump.DumpCommand)
	rootCmd.AddCommand(entries.EntriesCommand)
	rootCmd.AddCommand(config.ConfigCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
