package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "configgus",
	Short: "The tool to help you centralize all needed config files to backup them.",
	Long: `Configgus is a very simple tool to help you centralize and organize your config files in one place.
    It could be a git repo to easily backup all of them.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
