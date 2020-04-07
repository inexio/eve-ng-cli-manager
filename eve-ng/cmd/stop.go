package cmd

import (
	stopsubcommands "eve-ng-cli-manager/eve-ng/cmd/stopSubcommands"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops a component",
	Long: `Stops the given component within the provided lab`,
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
	rootCmd.AddCommand(stopCmd)
	stopCmd.AddCommand(stopsubcommands.StopNodeCmd)
	stopCmd.AddCommand(stopsubcommands.StopNodesCmd)
	stopCmd.PersistentFlags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")
}
