package createsubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// CreateLabCmd represents the createLab command
var CreateLabCmd = &cobra.Command{
	Use:   "lab",
	Short: "Creates a lab",
	Long:  `Creates a lab in the given directory`,
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
		version := cmd.Flag("version").Value.String()
		author := cmd.Flag("author").Value.String()
		description := cmd.Flag("description").Value.String()
		body := cmd.Flag("body").Value.String()

		//Perform AddLab operation
		err = client.AddLab(path, name, version, author, description, body)
		if err != nil {
			log.Error().
				Msg("Error during AddLab")
			os.Exit(1)
		}

		fmt.Println("Successfully created lab.")
	},
}

func init() {
	CreateLabCmd.Flags().StringP("path", "p", "", "The Path the lab will be created in")
	err := CreateLabCmd.MarkFlagRequired("path")
	if err != nil {
		log.Error().
			Msg("Could not mark 'path' flag required")
		os.Exit(1)
	}

	CreateLabCmd.Flags().StringP("name", "n", "", "Set the name of the lab")
	err = CreateLabCmd.MarkFlagRequired("name")
	if err != nil {
		log.Error().
			Msg("Could not mark 'name' flag required")
		os.Exit(1)
	}

	CreateLabCmd.Flags().StringP("version", "v", "", "Set the labs version")
	err = CreateLabCmd.MarkFlagRequired("version")
	if err != nil {
		log.Error().
			Msg("Could not mark 'version' flag required")
		os.Exit(1)
	}

	CreateLabCmd.Flags().StringP("author", "a", "", "Set the name of the author")
	err = CreateLabCmd.MarkFlagRequired("author")
	if err != nil {
		log.Error().
			Msg("Could not mark 'author' flag required")
		os.Exit(1)
	}

	CreateLabCmd.Flags().StringP("description", "d", "", "Set the labs description")
	err = CreateLabCmd.MarkFlagRequired("description")
	if err != nil {
		log.Error().
			Msg("Could not mark 'description' flag required")
		os.Exit(1)
	}

	CreateLabCmd.Flags().StringP("body", "b", "", "Set the labs body")
	err = CreateLabCmd.MarkFlagRequired("body")
	if err != nil {
		log.Error().
			Msg("Could not mark 'body' flag required")
		os.Exit(1)
	}
}
