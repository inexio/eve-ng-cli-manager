package cmd

import (
	editsubcommands "eve-ng-cli-manager/eve-ng/cmd/editSubcommands"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edits a component",
	Long:  `Edit the given component`,
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
	rootCmd.AddCommand(editCmd)
	editCmd.AddCommand(editsubcommands.EditLabCmd)
	editCmd.AddCommand(editsubcommands.EditUserCmd)
}
