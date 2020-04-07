package getsubcommands

import (
	"os"
	"strconv"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// GetNodeCmd represents the getNode command
var GetNodeCmd = &cobra.Command{
	Use:   "node <node-id>",
	Args:  cobra.ExactArgs(1),
	Short: "Returns data about a single node",
	Long: `Returns data about the given node within the provided lab

The data returned contains:
	- uuid
	- name
	- type
	- status
	- template
	- cpu
	- ram
	- image
	- console
	- ethernet
	- delay
	- icon
	- url
	- top
	- left
	- config
	- firstmac
	- configlist`,
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

		//Parse lab var
		var lab string
		if cmd.Flag("lab").Changed {
			lab = cmd.Flag("lab").Value.String()
		} else {
			lab = viper.GetString("labPath") + viper.GetString("labName") + ".unl"
		}

		//Parse node-id var
		nodeID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Error().
				Msg("Couldn't convert networkID to int")
			os.Exit(1)
		}

		//Perform GetNode operation
		node, err := client.GetNode(lab, nodeID)
		if err != nil {
			log.Error().
				Msg("Error during GetLabNode")
			os.Exit(1)
		}

		//Print the node object
		err = PrintData(format, node, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}

func init() {
	GetNodeCmd.Flags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")
}
