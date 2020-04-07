package disconnectsubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// NodeFromNetworkCmd represents the nodeFromNetworkCmd command
var NodeFromNetworkCmd = &cobra.Command{
	Use:   "node-from-network",
	Short: "Disconnects a node from a network",
	Long:  `Disconnects the given node via the given interface from a network inside the provided lab`,
	PreRun: func(cmd *cobra.Command, args []string) {
		lab := viper.GetString("labPath") + viper.GetString("labName")

		if lab == "" {
			err := cmd.MarkFlagRequired("lab")
			if err != nil {
				log.Error().
					Msg("Could not mark 'lab' flag required")
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
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

		//Parse lab var
		var lab string
		if cmd.Flag("lab").Changed {
			lab = cmd.Flag("lab").Value.String()
		} else {
			lab = viper.GetString("labPath") + viper.GetString("labName") + ".unl"
		}

		//Parse node-, network- and interface-id
		nodeID, err := cmd.Flags().GetInt("node")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'node' flag value")
			os.Exit(1)
		}
		interfaceId, err := cmd.Flags().GetInt("interface")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'interface' flag value")
			os.Exit(1)
		}

		//Perform DisconnectNodeInterfaceFromNetwork operation
		err = client.DisconnectNodeInterfaceFromNetwork(lab, nodeID, interfaceId)
		if err != nil {
			log.Error().
				Msg("Error during DisconnectNodeInterfaceFromNetwork")
			os.Exit(1)
		}

		fmt.Println("Successfully disconnect interface", interfaceId, "of node", nodeID, "from its network.")
	},
}

func init() {
	NodeFromNetworkCmd.Flags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")

	NodeFromNetworkCmd.Flags().Int("node", 0, "Set the id of the node")
	err := NodeFromNetworkCmd.MarkFlagRequired("node")
	if err != nil {
		log.Error().
			Msg("Could not mark 'node' flag required")
		os.Exit(1)
	}

	NodeFromNetworkCmd.Flags().Int("interface", 0, "Set the id of the node-interface")
	err = NodeFromNetworkCmd.MarkFlagRequired("interface")
	if err != nil {
		log.Error().
			Msg("Could not mark 'interface' flag required")
		os.Exit(1)
	}
}
