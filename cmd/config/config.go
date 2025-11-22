package config

import (
	"fmt"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tpyle/ksv/lib/cfg"
)

var ConfigCommand = &cobra.Command{
	Use:   "config",
	Short: "Configure Your Backend Storage",
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all configuration keys and values",
	Run: func(cmd *cobra.Command, args []string) {
		keys := viper.AllKeys()
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Printf("%s = %v\n", key, viper.Get(key))
		}
	},
}

var setCommand = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a configuration key to a specific value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		viper.Set(key, value)

		var config cfg.Config
		err := viper.Unmarshal(&config)
		if err != nil {
			logrus.WithError(err).Fatal("Invalid config")
			return
		}

		err = config.Validate()
		if err != nil {
			logrus.WithError(err).Fatal("Invalid config")
			return
		}

		if err := viper.WriteConfig(); err != nil {
			logrus.WithError(err).Fatal("Error writing config")
			return
		}

		fmt.Printf("%s = %s\n", key, value)
	},
}

var unsetCommand = &cobra.Command{
	Use:   "unset [key]",
	Short: "Unset a configuration key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		viper.Set(key, nil)

		var config cfg.Config
		err := viper.Unmarshal(&config)
		if err != nil {
			logrus.WithError(err).Fatal("Invalid config")
			return
		}

		err = config.Validate()
		if err != nil {
			logrus.WithError(err).Fatal("Invalid config")
			return
		}

		if err := viper.WriteConfig(); err != nil {
			logrus.WithError(err).Fatal("Error writing config")
			return
		}

		fmt.Printf("Unset %s\n", key)
	},
}

var getCommand = &cobra.Command{
	Use:   "get [key]",
	Short: "Get the value of a configuration key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: key is required unless --all is specified")
			return
		}

		key := args[0]
		value := viper.Get(key)
		if value == nil {
			fmt.Printf("Key %s not found\n", key)
			return
		}
		fmt.Printf("%v\n", value)
	},
}

func init() {
	ConfigCommand.AddCommand(listCommand)
	ConfigCommand.AddCommand(setCommand)
	ConfigCommand.AddCommand(unsetCommand)
	ConfigCommand.AddCommand(getCommand)
}
