package cmd

import (
	deletesubcommands "eve-ng-cli-manager/eve-ng/cmd/deleteSubcommands"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a component",
	Long:  `Deletes the given component within the given lab`,
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
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deletesubcommands.DeleteFolderCmd)
	deleteCmd.AddCommand(deletesubcommands.DeleteLabCmd)
	deleteCmd.AddCommand(deletesubcommands.DeleteNetworkCmd)
	deleteCmd.AddCommand(deletesubcommands.DeleteNodeCmd)
	deleteCmd.AddCommand(deletesubcommands.DeleteUserCmd)
}
