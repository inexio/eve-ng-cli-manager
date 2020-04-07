package cmd

import (
	createsubcommands "eve-ng-cli-manager/eve-ng/cmd/createSubcommands"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <sub-command>",
	Short: "Creates a component",
	Long:  `Use this command to create a certain component`,
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
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createsubcommands.CreateFolderCmd)
	createCmd.AddCommand(createsubcommands.CreateLabCmd)
	createCmd.AddCommand(createsubcommands.CreateNetworkCmd)
	createCmd.AddCommand(createsubcommands.CreateNodeCmd)
	createCmd.AddCommand(createsubcommands.CreateUserCmd)
}
