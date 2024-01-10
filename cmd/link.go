package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/kadencartwright/dotman/pkg/config"
)

func init() {
	rootCmd.AddCommand(linkCmd)

}

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "link config files from the config repo to their specified locations on disk",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgFile := cmd.Flag("file").Value.String()
		cfgDir := filepath.Dir(cfgFile)
		err := os.Chdir(cfgDir)
		if err != nil {
			return fmt.Errorf("could not cd into specified directory: ", cfgDir)
		}
		configData, err := config.NewConfigFromFile(cfgFile)
		if err != nil {
			return fmt.Errorf("could not find the given config file")
		}
		errs := configData.Link()

		if len(errs) != 0 {
			return fmt.Errorf("could not link files")
		}
		return nil
	},
}
