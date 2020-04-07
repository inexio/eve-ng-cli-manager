package cmd

import (
	startsubcommands "eve-ng-cli-manager/eve-ng/cmd/startSubcommands"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a component",
	Long: `Starts the given component within the provided lab`,
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
	rootCmd.AddCommand(startCmd)
	startCmd.AddCommand(startsubcommands.StartNodeCmd)
	startCmd.AddCommand(startsubcommands.StartNodesCmd)
	startCmd.PersistentFlags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")
}
