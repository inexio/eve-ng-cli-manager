package createsubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// CreateFolderCmd represents the createFolder command
var CreateFolderCmd = &cobra.Command{
	Use:   "folder",
	Short: "Creates a folder",
	Long:  `Creates a folder within the given path`,
	Args:  cobra.ExactArgs(0),
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
		path := cmd.Flag("path").Value.String()
		name := cmd.Flag("name").Value.String()

		//Perform AddFolder operation
		err = client.AddFolder(path, name)
		if err != nil {
			log.Error().
				Msg("Error during AddFolder")
			os.Exit(1)
		}

		fmt.Println("Successfully created folder.")
	},
}

func init() {
	CreateFolderCmd.Flags().String("path", "", "Set the path of the folder")
	err := CreateFolderCmd.MarkFlagRequired("path")
	if err != nil {
		log.Error().
			Msg("Could not mark 'path' flag required")
		os.Exit(1)
	}

	CreateFolderCmd.Flags().String("name", "", "Set the name of the folder")
	err = CreateFolderCmd.MarkFlagRequired("name")
	if err != nil {
		log.Error().
			Msg("Could not mark 'name' flag required")
		os.Exit(1)
	}
}
