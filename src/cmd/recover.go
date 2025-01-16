package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// recoverCmd represents the recover command
var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "Set the config file from the configgus main directory to their origin",
	Long:  `Allows to copy vonfig files from the configgus main directory to their origin.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("recover called")
	},
}

func init() {
	rootCmd.AddCommand(recoverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recoverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recoverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
