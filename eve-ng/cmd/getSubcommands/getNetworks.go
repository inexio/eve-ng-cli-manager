package getsubcommands

import (
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetNetworksCmd represents the getNetworks command
var GetNetworksCmd = &cobra.Command{
	Use:   "networks",
	Args:  cobra.ExactArgs(0),
	Short: "Returns a list of all networks",
	Long: `Returns a list of all networks within the provided lab

The data returned contains:
	- list of networks:
		- id
		- count
		- name
		- type
		- top
		- left
		- style
		- linkstyle
		- color
		- label
		- visibility`,
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

		//Perform GetNetworks operation
		networks, err := client.GetNetworks(lab)
		if err != nil {
			log.Error().
				Msg("Error during GetLabNetworks")
			os.Exit(1)
		}

		//Print the network objects
		err = PrintData(format, networks, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}

func init() {
	GetNetworksCmd.Flags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")
}
