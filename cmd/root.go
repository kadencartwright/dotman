package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dotman",
	Short: "dotman is a dotfiles management tool",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	var CfgFile string
	rootCmd.PersistentFlags().StringVarP(&CfgFile, "file", "f", "", "config file to use (required)")
	rootCmd.MarkPersistentFlagRequired("file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
