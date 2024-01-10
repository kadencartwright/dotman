package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/kadencartwright/dotman/pkg/config"
)

func init() {
	rootCmd.AddCommand(copyCmd)

}

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy config files from their specified locations on disk into the config repo",
	RunE: func(cmd *cobra.Command, args []string) error {
		flags := cmd.PersistentFlags()
		cfgFile, err := flags.GetString("file")
		if err != nil {
			return fmt.Errorf("no config file specified")

		}

		configData, err := config.NewConfigFromFile(cfgFile)
		if err != nil {
			return fmt.Errorf("could not find the given config file or none specified")

		}
		errs := configData.Copy()

		if len(errs) != 0 {
			for _, v := range errs {
				fmt.Println(v.Error())
			}
			return fmt.Errorf("could not copy files")
		}

		return nil
	},
}
