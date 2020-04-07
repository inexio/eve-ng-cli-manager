package deletesubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// DeleteUserCmd represents the deleteUser command
var DeleteUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Deletes a user",
	Long:  `Deletes a user via its username`,
	Args:  cobra.ExactArgs(1),
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

		//Parse username var
		username := args[0]

		//Perform RemoveUser operation
		err = client.RemoveUser(username)
		if err != nil {
			log.Error().
				Msg("Error during RemoveUser")
			os.Exit(1)
		}

		fmt.Println("Successfully deleted user", username+".")
	},
}
