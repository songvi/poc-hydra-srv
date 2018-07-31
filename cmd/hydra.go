package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Launch hydra - oauth2 server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := doNothing(); err != nil {
			log.Fatal(err)
		}
	},
}

func doNothing() error {
	fmt.Println("Just print out a text line")
	return nil
}

func init() {
	RootCmd.AddCommand(StartCmd)
}

/*
func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Cant' read config: ", err)
		os.Exit(1)
	}
}
*/
