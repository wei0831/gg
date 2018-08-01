package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wei0831/ghg/utli"
)

var cmdOpen = &cobra.Command{
	Use:   "open [path](default=\".\")",
	Short: "Launches the source remote URL for the target path with the given branch.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			utli.Open(".", remote, branch, false)
		} else {
			utli.Open(args[0], remote, branch, false)
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdOpen)
	cmdOpen.PersistentFlags().StringVarP(&remote, "remote", "r", "origin", "remote")
	cmdOpen.PersistentFlags().StringVarP(&branch, "branch", "b", "master", "branch")
}
