package cmd

import (
	wipesubcommands "eve-ng-cli-manager/eve-ng/cmd/wipeSubcommands"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// wipeCmd represents the wipe command
var wipeCmd = &cobra.Command{
	Use:   "wipe",
	Short: "Wipes a component",
	Long:  `Wipes the given component`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Error().
				Msg("Could not display help")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(wipeCmd)
	wipeCmd.AddCommand(wipesubcommands.WipeNodeCmd)
	wipeCmd.AddCommand(wipesubcommands.WipeNodesCmd)
	wipeCmd.PersistentFlags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")
}
