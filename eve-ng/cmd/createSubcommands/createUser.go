package createsubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// CreateUserCmd represents the createUser command
var CreateUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Creates a user",
	Long:  `Creates a new user`,
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
		username := cmd.Flag("username").Value.String()
		name := cmd.Flag("name").Value.String()
		email := cmd.Flag("email").Value.String()
		password := cmd.Flag("password").Value.String()
		role := cmd.Flag("role").Value.String()
		expiration := cmd.Flag("expiration").Value.String()
		dateStart := cmd.Flag("start").Value.String()
		extAuth := cmd.Flag("ext-auth").Value.String()
		pExpiration := cmd.Flag("pexpiration").Value.String()
		pod, err := cmd.Flags().GetInt("pod")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'pod' flag value")
			os.Exit(1)
		}
		cpu, err := cmd.Flags().GetInt("cpu")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'cpu' flag value")
			os.Exit(1)
		}
		ram, err := cmd.Flags().GetInt("ram")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'ram' flag value")
			os.Exit(1)
		}

		//Perform AddUser operation
		err = client.AddUser(username, name, email, password, role, expiration, dateStart, extAuth, pod, pExpiration, cpu, ram)
		if err != nil {
			log.Error().
				Msg("Error during AddUser")
			os.Exit(1)
		}

		fmt.Println("Successfully created new user", username+".")
	},
}

func init() {
	CreateUserCmd.Flags().String("username", "", "The username of the user")
	err := CreateUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Error().
			Msg("Could not mark 'username' flag required")
		os.Exit(1)
	}

	CreateUserCmd.Flags().String("name", "", "The name of the user")
	err = CreateUserCmd.MarkFlagRequired("name")
	if err != nil {
		log.Error().
			Msg("Could not mark 'name' flag required")
		os.Exit(1)
	}

	CreateUserCmd.Flags().String("password", "", "The password of the user")
	err = CreateUserCmd.MarkFlagRequired("password")
	if err != nil {
		log.Error().
			Msg("Could not mark 'password' flag required")
		os.Exit(1)
	}

	CreateUserCmd.Flags().Int("pod", 0, "The id of the user (has to be unique)")
	err = CreateUserCmd.MarkFlagRequired("pod")
	if err != nil {
		log.Error().
			Msg("Could not mark 'pod' flag required")
		os.Exit(1)
	}

	CreateUserCmd.Flags().String("email", "", "The email address of the user")

	CreateUserCmd.Flags().String("role", "admin", "The user-role of the user")

	CreateUserCmd.Flags().String("expiration", "-1", "The expiration date of the user (-1 means no expiration)")

	CreateUserCmd.Flags().String("start", "-1", "The start date of the user (-1 means start-date now)")

	CreateUserCmd.Flags().String("ext-auth", "internal", "The ext-auth method")

	CreateUserCmd.Flags().String("pexpiration", "-1", "The pexpiration date of the user (-1 means no pexpiration)")

	CreateUserCmd.Flags().Int("cpu", -1, "The cpu usage limit of the user (-1 means no limit)")

	CreateUserCmd.Flags().Int("ram", -1, "The ram usage limit of the user (-1 means no limit)")
}
