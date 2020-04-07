package cmd

import (
	exportsubcommands "eve-ng-cli-manager/eve-ng/cmd/exportSubcommands"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports a component",
	Long:  `Export the given component`,
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
	rootCmd.AddCommand(exportCmd)
	exportCmd.AddCommand(exportsubcommands.ExportNodeCmd)
	exportCmd.AddCommand(exportsubcommands.ExportNodesCmd)
	exportCmd.PersistentFlags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")
}
