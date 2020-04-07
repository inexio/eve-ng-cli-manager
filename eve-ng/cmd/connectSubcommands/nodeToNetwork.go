package connectsubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// NodeToNetworkCmd represents the nodeWithNetwork command
var NodeToNetworkCmd = &cobra.Command{
	Use:   "node-to-network",
	Short: "Connects a node to a network",
	Long:  `Connects a node via the given interface to a network inside the provided lab`,
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
		interfaceID, err := cmd.Flags().GetInt("interface")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'interface' flag value")
			os.Exit(1)
		}
		networkID, err := cmd.Flags().GetInt("network")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'network' flag value")
			os.Exit(1)
		}

		//Perform ConnectNodeInterfaceToNetwork operation
		err = client.ConnectNodeInterfaceToNetwork(lab, nodeID, interfaceID, networkID)
		if err != nil {
			log.Error().
				Msg("Error during ConnectNodeInterfaceToNetwork")
			os.Exit(1)
		}

		fmt.Println("Successfully connect interface", interfaceID, "of node", nodeID, "to network", networkID)
	},
}

func init() {
	NodeToNetworkCmd.Flags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")

	NodeToNetworkCmd.Flags().Int("node", 0, "Set the id of the node")
	err := NodeToNetworkCmd.MarkFlagRequired("node")
	if err != nil {
		log.Error().
			Msg("Could not mark 'node' flag required")
		os.Exit(1)
	}

	NodeToNetworkCmd.Flags().Int("interface", 0, "Set the id of the node-interface")
	err = NodeToNetworkCmd.MarkFlagRequired("interface")
	if err != nil {
		log.Error().
			Msg("Could not mark 'interface' flag required")
		os.Exit(1)
	}

	NodeToNetworkCmd.Flags().Int("network", 0, "Set the id of the network")
	err = NodeToNetworkCmd.MarkFlagRequired("network")
	if err != nil {
		log.Error().
			Msg("Could not mark 'network' flag required")
		os.Exit(1)
	}
}
