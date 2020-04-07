package getsubcommands

import (
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetLabFilesCmd represents the getLabFiles command
var GetLabFilesCmd = &cobra.Command{
	Use:   "lab-files",
	Args:  cobra.ExactArgs(1),
	Short: "Returns all lab files within a folder",
	Long: `Returns a list of all lab files contained in the given folder.

The name of the lab as well as the filename is displayed`,
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

		//Parse path var
		path := args[0]

		//Perform GetLabFiles operation
		labFiles, err := client.GetLabFiles(path)
		if err != nil {
			log.Error().
				Msg("Error during GetLabFiles")
			os.Exit(1)
		}

		//Print the lab-files object
		err = PrintData(format, labFiles, prettified, depth)
		if err != nil {
			log.Error().
				Msg("Could not print data correctly")
			os.Exit(1)
		}
	},
}
