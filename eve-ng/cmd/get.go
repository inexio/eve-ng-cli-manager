package cmd

import (
	getsubcommands "eve-ng-cli-manager/eve-ng/cmd/getSubcommands"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns data of a component",
	Long: `Returns data of a given component in either a human-readable, xml or json format.

If no specific format is given the default is set to 'human-readable'`,
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
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringP("format", "f", "human-readable", "Use this flag to set the output format")
	getCmd.PersistentFlags().StringP("depth", "d", "3", "Use this flag to set the depth of human-readable output")
	getCmd.PersistentFlags().BoolP("pretty", "p", false, "Use this flag to prettify the output if xml or json is set as output format")
	getCmd.AddCommand(getsubcommands.GetFoldersCmd)
	getCmd.AddCommand(getsubcommands.GetLabCmd)
	getCmd.AddCommand(getsubcommands.GetLabFilesCmd)
	getCmd.AddCommand(getsubcommands.GetNetworkCmd)
	getCmd.AddCommand(getsubcommands.GetNetworksCmd)
	getCmd.AddCommand(getsubcommands.GetNetworkTypesCmd)
	getCmd.AddCommand(getsubcommands.GetNodeCmd)
	getCmd.AddCommand(getsubcommands.GetNodeInterfacesCmd)
	getCmd.AddCommand(getsubcommands.GetNodesCmd)
	getCmd.AddCommand(getsubcommands.GetNodeTemplateCmd)
	getCmd.AddCommand(getsubcommands.GetNodeTemplatesCmd)
	getCmd.AddCommand(getsubcommands.GetSystemStatusCmd)
	getCmd.AddCommand(getsubcommands.GetTopologyCmd)
	getCmd.AddCommand(getsubcommands.GetUserCmd)
	getCmd.AddCommand(getsubcommands.GetUserRolesCmd)
	getCmd.AddCommand(getsubcommands.GetUsersCmd)
}
