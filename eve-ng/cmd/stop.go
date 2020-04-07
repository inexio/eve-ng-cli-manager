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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
