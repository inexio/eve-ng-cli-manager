package getsubcommands

import (
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetNodeTemplatesCmd represents the getNodeTemplates command
var GetNodeTemplatesCmd = &cobra.Command{
	Use:   "node-templates",
	Args:  cobra.ExactArgs(0),
	Short: "Returns a list of all node templates",
	Long: `Returns a list of all node templates configured on the EveNG server

The list consists of the template names and a brief description:

	template-name:description`,
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

		//Perform GetNodeTemplates operation
		nodeTemplates, err := client.GetNodeTemplates()
		if err != nil {
			log.Error().
				Msg("Error during GetNodeTemplates")
			os.Exit(1)
		}

		//Print node-templates object
		err = PrintData(format, nodeTemplates, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}
