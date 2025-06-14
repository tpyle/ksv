package cmd

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tpyle/ksv/cmd/config"
	"github.com/tpyle/ksv/cmd/dump"
	"github.com/tpyle/ksv/cmd/entries"
	"github.com/tpyle/ksv/cmd/search"
	"github.com/tpyle/ksv/cmd/sites"
	"github.com/tpyle/ksv/cmd/util"
	"github.com/tpyle/ksv/lib/cfg"
)

var (
	verbosity  int
	configFile string
	namespace  string
)

func getDefaultConfigFile() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("AppData"), "ksv", "config.yaml")
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "ksv", "config.yaml")
	default: // "linux" and other Unix-like systems
		return filepath.Join(os.Getenv("HOME"), ".config", "ksv", "config.yaml")
	}
}

var rootCmd = &cobra.Command{
	Use:   "ksv",
	Short: "KSV is a CLI tool for managing secrets",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set the logging level based on the verbosity flag
		switch verbosity {
		case 1:
			logrus.SetLevel(logrus.InfoLevel)
		case 2:
			logrus.SetLevel(logrus.DebugLevel)
		case 3:
			logrus.SetLevel(logrus.TraceLevel)
		default:
			logrus.SetLevel(logrus.WarnLevel)
		}

		config, err := cfg.LoadConfig(configFile)
		if err != nil {
			logrus.Fatalf("Error loading config: %v", err)
		}

		cmd.SetContext(util.AttachNamespaceToContext(namespace, cmd.Context()))
		cmd.SetContext(util.AttachConfigToContext(config, cmd.Context()))

		logrus.Tracef("Using Config %+v", util.GetConfigFromContext(cmd.Context()))
	},
}

func init() {
	rootCmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "Set the verbosity level")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config-file", "c", getDefaultConfigFile(), "Path to the config file")

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
