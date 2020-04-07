package editsubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// EditLabCmd represents the editLab command
var EditLabCmd = &cobra.Command{
	Use:   "lab",
	Short: "Edits a lab",
	Long:  `Edits the provided lab to change previously set values`,
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

		//Parse vars
		var lab string
		if cmd.Flag("lab").Changed {
			lab = cmd.Flag("lab").Value.String()
		} else {
			lab = viper.GetString("labPath") + viper.GetString("labName") + ".unl"
		}
		name := cmd.Flag("name").Value.String()
		version := cmd.Flag("version").Value.String()
		author := cmd.Flag("author").Value.String()
		description := cmd.Flag("description").Value.String()

		//Perform EditLab operation
		err = client.EditLab(lab, name, version, author, description)
		if err != nil {
			log.Error().
				Msg("Error during EditLab")
			os.Exit(1)
		}

		if !cmd.Flag("lab").Changed && cmd.Flag("name").Changed {
			viper.Set("labName", name)
			viper.WriteConfig()
		}

		fmt.Println("Successfully edited lab.")
	},
}

func init() {
	EditLabCmd.Flags().String("lab", "", "Set the lab-file path (only necessary if not already set via config)")

	EditLabCmd.Flags().StringP("name", "n", "", "Set the name of the lab")

	EditLabCmd.Flags().StringP("version", "v", "", "Set the version of the lab")

	EditLabCmd.Flags().StringP("author", "a", "", "Set the author of the lab")

	EditLabCmd.Flags().StringP("description", "d", "", "Set the description of the lab")
}
