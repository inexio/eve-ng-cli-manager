package deletesubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// DeleteFolderCmd represents the deleteFolder command
var DeleteFolderCmd = &cobra.Command{
	Use:   "folder <path/to/the/folder>",
	Short: "Deletes a folder",
	Long:  `Deletes the folder that resembles the given path`,
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

		//Parse path var
		path := args[0]

		//Perform RemoveFolder operation
		err = client.RemoveFolder(path)
		if err != nil {
			log.Error().
				Msg("Error during RemoveFolder")
			os.Exit(1)
		}

		fmt.Println("Successfully deleted folder.")
	},
}
