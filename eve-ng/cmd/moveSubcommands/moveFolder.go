package movesubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// MoveFolderCmd represents the moveFolder command
var MoveFolderCmd = &cobra.Command{
	Use:   "folder",
	Short: "Moves a folder",
	Long: `Move the given folder to the provided destination path
	
In the folders destination the folders name also has to be given.

Example: move folder --folder "/Testfolder" --destination "/Testfolder2/Testfolder"
	`,
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

		//Parse folder and destination vars
		folder := cmd.Flag("folder").Value.String()
		destination := cmd.Flag("destination").Value.String()

		//Perform MoveLab operation
		err = client.MoveFolder(folder, destination)
		if err != nil {
			log.Error().
				Msg("Error during MoveFolder")
			os.Exit(1)
		}

		fmt.Println("Successfully moved folder to new location.")
	},
}

func init() {
	MoveFolderCmd.Flags().String("folder", "", "Set the folders path")
	err := MoveFolderCmd.MarkFlagRequired("folder")
	if err != nil {
		log.Error().
			Msg("Could not mark 'folder' flag required")
		os.Exit(1)
	}

	MoveFolderCmd.Flags().String("destination", "", "Set the destination")
	err = MoveFolderCmd.MarkFlagRequired("destination")
	if err != nil {
		log.Error().
			Msg("Could not mark 'destination' flag required")
		os.Exit(1)
	}
}
