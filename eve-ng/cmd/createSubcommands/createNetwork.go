package createsubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// CreateNetworkCmd represents the createNetwork command
var CreateNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Creates a network",
	Long:  `Creates a network in the provided lab`,
	Args:  cobra.ExactArgs(0),
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

		//Parse vars
		var lab string
		if cmd.Flag("lab").Changed {
			lab = cmd.Flag("lab").Value.String()
		} else {
			lab = viper.GetString("labPath") + viper.GetString("labName") + ".unl"
		}
		networkType := cmd.Flag("type").Value.String()
		name := cmd.Flag("name").Value.String()
		left, err := cmd.Flags().GetInt("left")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'left' flag value")
			os.Exit(1)
		}
		top, err := cmd.Flags().GetInt("top")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'top' flag value")
			os.Exit(1)
		}
		visibility, err := cmd.Flags().GetInt("visibility")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'visibility' flag value")
			os.Exit(1)
		}
		postfix, err := cmd.Flags().GetInt("postfix")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'postfix' flag value")
			os.Exit(1)
		}

		//Perform AddNetwork operation
		networkId, err := client.AddNetwork(lab, networkType, name, left, top, visibility, postfix)
		if err != nil {
			log.Error().
				Msg("Error during AddNetwork")
			os.Exit(1)
		}

		fmt.Println("Successfully added network to lab.")
		fmt.Println("Network ID:", networkId)
	},
}

func init() {
	CreateNetworkCmd.Flags().StringP("lab", "l", "", "Set the lab-file path (only necessary if not already set via config)")
	CreateNetworkCmd.Flags().StringP("type", "t", "", "Set the network type")
	err := CreateNetworkCmd.MarkFlagRequired("type")
	if err != nil {
		log.Error().
			Msg("Could not mark 'type' flag required")
		os.Exit(1)
	}

	CreateNetworkCmd.Flags().StringP("name", "n", "", "Set the name of the network")
	err = CreateNetworkCmd.MarkFlagRequired("name")
	if err != nil {
		log.Error().
			Msg("Could not mark 'name' flag required")
		os.Exit(1)
	}

	CreateNetworkCmd.Flags().Int("left", 0, "Set the left value")
	err = CreateNetworkCmd.MarkFlagRequired("left")
	if err != nil {
		log.Error().
			Msg("Could not mark 'left' flag required")
		os.Exit(1)
	}

	CreateNetworkCmd.Flags().Int("top", 0, "Set the top value")
	err = CreateNetworkCmd.MarkFlagRequired("top")
	if err != nil {
		log.Error().
			Msg("Could not mark 'top' flag required")
		os.Exit(1)
	}

	CreateNetworkCmd.Flags().IntP("visibility", "v", 1, "Set the networks visibility")

	CreateNetworkCmd.Flags().IntP("postfix", "p", 0, "Set the network postfix")
}
