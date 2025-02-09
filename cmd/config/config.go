package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ConfigCommand = &cobra.Command{
	Use:   "config",
	Short: "Configure Your Backend Storage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}
