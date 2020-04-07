package getsubcommands

import (
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetUserCmd represents the getUser command
var GetUserCmd = &cobra.Command{
	Use:   "user <username>",
	Args:  cobra.ExactArgs(1),
	Short: "Returns data about a user",
	Long: `Returns detailed information about the given user

The returned topology contains information about:
	- destination
	- source
	- type
	- networkID`,
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

		//Parse username var
		username := args[0]

		//Perform GetUser operation
		user, err := client.GetUser(username)
		if err != nil {
			log.Error().
				Msg("Error during GetUser")
			os.Exit(1)
		}

		//Print user object
		err = PrintData(format, user, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}
