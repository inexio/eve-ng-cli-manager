package getsubcommands

import (
	"os"
	"strconv"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// GetNodeInterfacesCmd represents the getNodeInterface command
var GetNodeInterfacesCmd = &cobra.Command{
	Use:   "node-interfaces <node-id>",
	Args:  cobra.ExactArgs(1),
	Short: "Returns information about a nodes interfaces",
	Long: `Returns a list of all interfaces of the given node

The returned list contains further data on each interface`,
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

		//Perform GetNodeInterfaces opeation
		nodeInterfaces, err := client.GetNodeInterfaces(lab, nodeID)
		if err != nil {
			log.Error().
				Msg("Error during GetLabNodeInterfaces")
			os.Exit(1)
		}

		//Print node-interfaces object
		err = PrintData(format, nodeInterfaces, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}

func init() {
	GetNodeInterfacesCmd.Flags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")
}
