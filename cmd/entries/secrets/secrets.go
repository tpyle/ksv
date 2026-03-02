package secrets

import "github.com/spf13/cobra"

var SecretsCommand = &cobra.Command{
	Use:   "secrets",
	Short: "Manage secrets",
}
