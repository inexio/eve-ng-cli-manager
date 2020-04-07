package cmd

import (
	connectsubcommands "eve-ng-cli-manager/eve-ng/cmd/connectSubcommands"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connects two components with each other",
	Long: `Used to connect two components with each other.

At the moment the only opperation that is suppoerted is adding a nodeInterface to a network`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {

		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.AddCommand(connectsubcommands.NodeToNetworkCmd)
}
