package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Hydra service - Oauth 2",
	Long:  "Hydra service - Oauth2 service",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.StringVarP(&configFile, "config", "c", "config.json", "Path to config file")
}
