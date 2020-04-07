package getsubcommands

import (
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetSystemStatusCmd represents the getSystemStatus command
var GetSystemStatusCmd = &cobra.Command{
	Use:   "system-status",
	Args:  cobra.ExactArgs(0),
	Short: "Returns data concerning the system health",
	Long: `Returns detailed information about the systems health.

The returned data-set contains:
	- Cached	cached data
	- Cpu		cpu load
	- Disk		disk usage
	- Dynamips	
	- Iol	
	- Mem		memory usage
	- Qemu		qemu status
	- Qemuversion	
	- Swap		
	- Version	eve ng software version`,
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

		//Perform GetSystemStatus operation
		systemStatus, err := client.GetSystemStatus()
		if err != nil {
			log.Error().
				Msg("Error during GetSystemStatus")
			os.Exit(1)
		}

		//Print system-status object
		err = PrintData(format, systemStatus, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}
