package deletesubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// DeleteLabCmd represents the deleteLab command
var DeleteLabCmd = &cobra.Command{
	Use:   "lab",
	Short: "Deletes a lab",
	Long: `Deletes the provided lab.

!!!Before deleting a lab make sure to stop all nodes on every registered user!!!`,
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

		//perform RemoveLab operation
		lab := cmd.Flag("lab").Value.String()

		err = client.RemoveLab(lab)
		if err != nil {
			log.Error().
				Msg("Error during RemoveLab")
			os.Exit(1)
		}

		fmt.Println("Successfully deleted lab.")
	},
}

func init() {
	DeleteLabCmd.Flags().String("lab", "", "Set the lab-file path")
	err := DeleteLabCmd.MarkFlagRequired("lab")
	if err != nil {
		log.Error().
			Msg("Could not mark 'lab' flag required")
		os.Exit(1)
	}
}
