package cmd

import (
	"errors"

	"github.com/orgs/storyui/goap/internal/exec"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use: "exec",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please select a environment")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		environment := args[0]
		d := exec.NewExec()
		d.Run(environment, goatConfig)
	},
}
