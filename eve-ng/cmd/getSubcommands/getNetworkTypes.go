package getsubcommands

import (
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetNetworkTypesCmd represents the getNetworkTypes command
var GetNetworkTypesCmd = &cobra.Command{
	Use:   "network-types",
	Args:  cobra.ExactArgs(0),
	Short: "Returns all network-types",
	Long:  `Returns a list of all possible network-types`,
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

		//Perform GetNetworkTypes operation
		networkTypes, err := client.GetNetworkTypes()
		if err != nil {
			log.Error().
				Msg("Error during GetLabNetworks")
			os.Exit(1)
		}

		//Print the network-types object
		err = PrintData(format, networkTypes, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}
