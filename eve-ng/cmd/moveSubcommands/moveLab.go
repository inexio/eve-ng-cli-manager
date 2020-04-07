package movesubcommands

import (
	"fmt"
	"os"
	"strings"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// MoveLabCmd represents the moveLab command
var MoveLabCmd = &cobra.Command{
	Use:   "lab",
	Short: "Moves a lab",
	Long:  `Moves the given lab to the provided destination path`,
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

		//Parse lab and path var
		var lab string
		if cmd.Flag("lab").Changed {
			lab = cmd.Flag("lab").Value.String()
		} else {
			lab = viper.GetString("labPath") + viper.GetString("labName") + ".unl"
		}
		path := cmd.Flag("path").Value.String()

		//Perform MoveLab operation
		err = client.MoveLab(lab, path)
		if err != nil {
			log.Error().
				Msg("Error during MoveLab")
			os.Exit(1)
		}

		//Change configured standard lab
		if !cmd.Flag("lab").Changed {
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			viper.Set("labPath", path)
			err = viper.WriteConfig()
			if err != nil {
				log.Error().
					Msg("Couldn't write config")
				os.Exit(1)
			}
		}

		fmt.Println("Successfully moved lab to new location.")
	},
}

func init() {
	MoveLabCmd.Flags().String("lab", "", "Set the lab-file path")

	MoveLabCmd.Flags().String("path", "", "Set the path")
	err := MoveLabCmd.MarkFlagRequired("path")
	if err != nil {
		log.Error().
			Msg("Could not mark 'path' flag required")
		os.Exit(1)
	}
}
