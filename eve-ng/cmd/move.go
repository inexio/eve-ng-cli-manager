package cmd

import (
	movesubcommands "eve-ng-cli-manager/eve-ng/cmd/moveSubcommands"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Moves a component",
	Long: `Moves the given component.

Currently only supports labs and folders`,
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
	rootCmd.AddCommand(moveCmd)
	moveCmd.AddCommand(movesubcommands.MoveFolderCmd)
	moveCmd.AddCommand(movesubcommands.MoveLabCmd)
}
