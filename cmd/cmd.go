package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/storyui/goat/config"
)

var rootCmd = &cobra.Command{
	Use: "goat",
}

var workDir string

var goatConfig = config.GoatConfig{}

func init() {
	rootCmd.PersistentFlags().StringVar(&workDir, "work-dir", ".", "working directory")
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(execCmd)
}

func Execute() {
	rootCmd.Execute()
}

func initConfig() {
	viper.SetConfigName("goat")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)
	log.Println(workDir)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err := viper.Unmarshal(&goatConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
