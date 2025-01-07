package cmd

import (
	"log"

	"github.com/orgs/storyui/goap/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "goat",
}

var goatConfig = config.GoatConfig{}

func init() {
	viper.SetConfigName("goat")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("playground")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err := viper.Unmarshal(&goatConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	rootCmd.AddCommand(execCmd)
}

func Execute() {
	rootCmd.Execute()
}
