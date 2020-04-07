package editsubcommands

import (
	"fmt"
	"os"
	"strconv"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// EditUserCmd represents the editUser command
var EditUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Edits a user",
	Long:  `Edits a user via the given username`,
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
		username := cmd.Flag("username").Value.String()
		name := cmd.Flag("name").Value.String()
		email := cmd.Flag("email").Value.String()
		password := cmd.Flag("password").Value.String()
		role := cmd.Flag("role").Value.String()
		expiration := cmd.Flag("expiration").Value.String()
		pExpiration := cmd.Flag("pexpiration").Value.String()
		var pod int
		podString := cmd.Flag("pod").Value.String()
		if podString != "" {
			pod, err = strconv.Atoi(podString)
			if err != nil {
				log.Error().
					Msg("Couldn't convert pod to int")
				os.Exit(1)
			}
		} else {
			userData, err := client.GetUser(username)
			if err != nil {
				log.Error().
					Msg("Error during GetUser")
				os.Exit(1)
			}

			pod, err = strconv.Atoi(userData.Pod)
			if err != nil {
				log.Error().
					Msg("Couldn't convert pod to int")
				os.Exit(1)
			}
		}

		//Perform AddUser operation
		err = client.EditUser(username, name, email, password, role, expiration, pod, pExpiration)
		if err != nil {
			log.Error().
				Msg("Error during AddUser")
			os.Exit(1)
		}

		fmt.Println("Successfully edited user", username+".")
	},
}

func init() {
	EditUserCmd.Flags().String("username", "", "The username of the user")
	err := EditUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Error().
			Msg("Could not mark 'username' flag required")
		os.Exit(1)
	}

	EditUserCmd.Flags().String("name", "", "The name of the user")

	EditUserCmd.Flags().String("email", "", "The email address of the user")

	EditUserCmd.Flags().String("password", "", "The password of the user")

	EditUserCmd.Flags().String("role", "admin", "The user-role of the user")

	EditUserCmd.Flags().String("expiration", "-1", "The expiration date of the user (-1 means no expiration)")

	EditUserCmd.Flags().String("pod", "", "The id of the user (has to be unique)")

	EditUserCmd.Flags().String("pexpiration", "-1", "The pexpiration date of the user (-1 means no pexpiration)")
}
