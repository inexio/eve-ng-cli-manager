package getsubcommands

import (
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetNodeTemplateCmd represents the getNodeTemplate command
var GetNodeTemplateCmd = &cobra.Command{
	Use:   "node-template <template-name>",
	Args:  cobra.ExactArgs(1),
	Short: "Returns data about a single node-template",
	Long: `Returns detailed information about the given node-template

Node templates are referenced by name`,
	Run: func(cmd *cobra.Command, args []string) {
		//Parse format, depth and prettified
		format, depth, prettified := parsePersistentFlags(cmd)

		//Create new EveNG client
		client, err := evengclient.NewEveNgClient(viper.GetString("baseUrl"))
		if err != nil {
			log.Error().
				Msg("Error during initialization of eve client")
			os.Exit(1)
		}

		err = client.SetUsernameAndPassword(viper.GetString("username"), viper.GetString("password"))
		if err != nil {
			log.Error().
				Msg("Error during SetUsernameAndPassword")
			os.Exit(1)
		}

		err = client.Login()
		if err != nil {
			log.Error().
				Msg("Error during login")
			os.Exit(1)
		}
		defer func() {
			err = client.Logout()
			if err != nil {
				log.Error().
					Msg("Error during logout")
				os.Exit(1)
			}
		}()

		//Parse template var
		template := args[0]

		//Perform GetNodeTemplate operation
		nodeTemplates, err := client.GetNodeTemplate(template)
		if err != nil {
			log.Error().
				Msg("Error during GetNodeTemplate")
			os.Exit(1)
		}

		//Print node-template object
		err = PrintData(format, nodeTemplates, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}
