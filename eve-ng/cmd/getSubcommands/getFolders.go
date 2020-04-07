package getsubcommands

import (
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetFoldersCmd represents the getFolders command
var GetFoldersCmd = &cobra.Command{
	Use:   "folders <path>",
	Args:  cobra.ExactArgs(1),
	Short: "Returns all folders within a folder",
	Long: `Returns a list of all folders contained within the given folder

Gives information about the folders name and its path`,
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

		//Parse var
		path := args[0]

		//perform GetFolders operation
		folders, err := client.GetFolders(path)
		if err != nil {
			log.Error().
				Msg("Error during GetFolders")
			os.Exit(1)
		}

		//Print the folders object
		err = PrintData(format, folders, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}
