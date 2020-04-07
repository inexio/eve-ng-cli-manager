package getsubcommands

import (
	"os"
	"strconv"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// GetNetworkCmd represents the getNetwork command
var GetNetworkCmd = &cobra.Command{
	Use:   "network <network-id>",
	Args:  cobra.ExactArgs(1),
	Short: "Returns data about a single network",
	Long: `Returns data about the given network within the provided lab

The data returned contains:
	- count
	- name
	- type
	- top
	- left
	- style
	- linkstyle
	- color
	- label
	- visibility
`,
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

		//Parse network-id
		networkId, err := strconv.Atoi(args[0])
		if err != nil {
			log.Error().
				Msg("Couldn't convert networkId to int")
			os.Exit(1)
		}

		//Perform GetNetwork operation
		network, err := client.GetNetwork(lab, networkId)
		if err != nil {
			log.Error().
				Msg("Error during GetLabNetworks")
			os.Exit(1)
		}

		//Print the network object
		err = PrintData(format, network, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}

func init() {
	GetNetworkCmd.Flags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")
}
