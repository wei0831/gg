package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wei0831/ghg/utli"
)

var cmdBlame = &cobra.Command{
	Use:   "blame [path](default=\".\")",
	Short: "Launches the blame remote URL for the target path with the given branch.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utli.Open(".", remote, branch, true)
		} else {
			utli.Open(args[0], remote, branch, true)
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdBlame)
	cmdBlame.PersistentFlags().StringVarP(&remote, "remote", "r", "origin", "remote")
	cmdBlame.PersistentFlags().StringVarP(&branch, "branch", "b", "master", "branch")
}
