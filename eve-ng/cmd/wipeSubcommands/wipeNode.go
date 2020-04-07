package wipesubcommands

import (
	"fmt"
	"os"
	"strconv"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// WipeNodeCmd represents the wipeNode command
var WipeNodeCmd = &cobra.Command{
	Use:   "node <node-id>",
	Args:  cobra.ExactArgs(1),
	Short: "Wipes a node",
	Long:  `Wipes the startup-configuration of the given node within the provided lab`,
	PreRun: func(cmd *cobra.Command, args []string) {
		lab := viper.GetString("labPath") + viper.GetString("labName")
		if lab == "" {
			err := cmd.MarkPersistentFlagRequired("lab")
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

		//Parse lab and node-id var
		var lab string
		if cmd.Flag("lab").Changed {
			lab = cmd.Flag("lab").Value.String()
		} else {
			lab = viper.GetString("labPath") + viper.GetString("labName") + ".unl"
		}
		nodeID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Error().
				Msg("Couldn't convert networkID to int")
			os.Exit(1)
		}

		//Perform WipeNode operation
		err = client.WipeNode(lab, nodeID)
		if err != nil {
			log.Error().
				Msg("Error during WipeNode")
			os.Exit(1)
		}

		fmt.Println("Node", nodeID, "has been wiped successfully.")
	},
}
