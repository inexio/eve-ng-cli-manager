package deletesubcommands

import (
	"fmt"
	"os"
	"strconv"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// DeleteNodeCmd represents the deleteNode command
var DeleteNodeCmd = &cobra.Command{
	Use:   "node <node-id>",
	Short: "Deletes a node",
	Long:  `Deletes a node via the given id in the provided lab`,
	Args:  cobra.ExactArgs(1),
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

		//perform RemoveLab operation
		var lab string
		if cmd.Flag("lab").Changed {
			lab = cmd.Flag("lab").Value.String()
		} else {
			lab = viper.GetString("labPath") + viper.GetString("labName") + ".unl"
		}
		nodeId, err := strconv.Atoi(args[0])
		if err != nil {
			log.Error().
				Msg("Couldn't convert node-id to int")
			os.Exit(1)
		}

		err = client.RemoveNode(lab, nodeId)
		if err != nil {
			log.Error().
				Msg("Error during Removenode")
			os.Exit(1)
		}

		fmt.Println("Successfully deleted node", strconv.Itoa(nodeId)+".")
	},
}

func init() {
	DeleteNodeCmd.Flags().String("lab", "", "Set the lab-file path")
}
