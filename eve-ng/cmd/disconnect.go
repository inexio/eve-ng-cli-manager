package cmd

import (
	disconnectsubcommands "eve-ng-cli-manager/eve-ng/cmd/disconnectSubcommands"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnects two components from each other",
	Long: `Disconnects the given component from the other given component.

At the moment the only opperation that is suppoerted is adding a nodeInterface to a network`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {

		}
	},
}

func init() {
	rootCmd.AddCommand(disconnectCmd)
	disconnectCmd.AddCommand(disconnectsubcommands.NodeFromNetworkCmd)
}
