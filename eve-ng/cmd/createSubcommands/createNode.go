package createsubcommands

import (
	"fmt"
	"os"

	evengclient "github.com/inexio/eve-ng-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// CreateNodeCmd represents the createNode command
var CreateNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Creates a node",
	Long:  `Creates a node in the provided lab`,
	Args:  cobra.ExactArgs(0),
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
		nodeType := cmd.Flag("type").Value.String()
		name := cmd.Flag("name").Value.String()
		config := cmd.Flag("config").Value.String()
		template := cmd.Flag("template").Value.String()
		icon := cmd.Flag("icon").Value.String()
		image := cmd.Flag("image").Value.String()
		console := cmd.Flag("console").Value.String()
		cpuLimit := cmd.Flag("cpu-limit").Value.String()
		firstMac := cmd.Flag("firstmac").Value.String()
		rdpUser := cmd.Flag("rdp-user").Value.String()
		rdpPass := cmd.Flag("rdp-pass").Value.String()
		uuid := cmd.Flag("uuid").Value.String()

		left, err := cmd.Flags().GetInt("left")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'left' flag value")
			os.Exit(1)
		}
		top, err := cmd.Flags().GetInt("top")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'top' flag value")
			os.Exit(1)
		}
		ram, err := cmd.Flags().GetInt("ram")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'ram' flag value")
			os.Exit(1)
		}
		delay, err := cmd.Flags().GetInt("delay")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'delay' flag value")
			os.Exit(1)
		}
		cpu, err := cmd.Flags().GetInt("cpu")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'cpu' flag value")
			os.Exit(1)
		}
		ethernet, err := cmd.Flags().GetInt("ethernet")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'ethernet' flag value")
			os.Exit(1)
		}
		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			log.Error().
				Msg("Couldn't read 'count' flag value")
			os.Exit(1)
		}

		//Perform AddNode operation
		nodeID, err := client.AddNode(lab, nodeType, template, config, delay, icon, image, name, left, top, ram, console, cpu, cpuLimit, ethernet, firstMac, rdpUser, rdpPass, uuid, count)
		if err != nil {
			log.Error().
				Msg("Error during AddNode")
			os.Exit(1)
		}

		fmt.Println("Successfully added node to lab.")
		fmt.Println("Node ID:", nodeID)
	},
}

func init() {
	CreateNodeCmd.Flags().StringP("type", "t", "", "Set the type of the node")
	err := CreateNodeCmd.MarkFlagRequired("type")
	if err != nil {
		log.Error().
			Msg("Could not mark 'type' flag required")
		os.Exit(1)
	}

	CreateNodeCmd.Flags().String("template", "", "Set the template that will be used")
	err = CreateNodeCmd.MarkFlagRequired("template")
	if err != nil {
		log.Error().
			Msg("Could not mark 'template' flag required")
		os.Exit(1)
	}

	CreateNodeCmd.Flags().StringP("name", "n", "", "Set the name of the node")
	err = CreateNodeCmd.MarkFlagRequired("name")
	if err != nil {
		log.Error().
			Msg("Could not mark 'name' flag required")
		os.Exit(1)
	}

	CreateNodeCmd.Flags().Int("left", 0, "Set the distance to the left side")
	err = CreateNodeCmd.MarkFlagRequired("left")
	if err != nil {
		log.Error().
			Msg("Could not mark 'left' flag required")
		os.Exit(1)
	}

	CreateNodeCmd.Flags().Int("top", 0, "Set the distance to the top")
	err = CreateNodeCmd.MarkFlagRequired("top")
	if err != nil {
		log.Error().
			Msg("Could not mark 'top' flag required")
		os.Exit(1)
	}

	CreateNodeCmd.Flags().Int("ram", 0, "Set the amount of ram dedicated to the node")
	err = CreateNodeCmd.MarkFlagRequired("ram")
	if err != nil {
		log.Error().
			Msg("Could not mark 'ram' flag required")
		os.Exit(1)
	}

	CreateNodeCmd.Flags().String("icon", "", "Set the icon that will be used for the node")
	err = CreateNodeCmd.MarkFlagRequired("icon")
	if err != nil {
		log.Error().
			Msg("Could not mark 'icon' flag required")
		os.Exit(1)
	}

	CreateNodeCmd.Flags().String("image", "", "Set the image that will be applied to the node")
	err = CreateNodeCmd.MarkFlagRequired("image")
	if err != nil {
		log.Error().
			Msg("Could not mark 'image' flag required")
		os.Exit(1)
	}

	CreateNodeCmd.Flags().StringP("lab", "l", "", "Set the lab-file path (only necessary if not already set via config)")

	CreateNodeCmd.Flags().StringP("config", "c", "0", "Set config of the node")

	CreateNodeCmd.Flags().IntP("delay", "d", 0, "Set the delay that will be applied while adding the Node")

	CreateNodeCmd.Flags().Int("count", 1, "Set the number of nodes that will be added")

	CreateNodeCmd.Flags().String("console", "telnet", "Set the console type of the node")

	CreateNodeCmd.Flags().Int("cpu", 1, "Set the amount of cpu that will be applied to the node")

	CreateNodeCmd.Flags().String("cpu-limit", "undefined", "Set the cpu limit for the node")

	CreateNodeCmd.Flags().Int("ethernet", 4, "Set the amount of ethernet-ports of the node")

	CreateNodeCmd.Flags().String("firstmac", "", "Set the firstmac of the node")

	CreateNodeCmd.Flags().String("rdp-user", "", "Set the rdp-user")

	CreateNodeCmd.Flags().String("rdp-pass", "", "Set the rdp-pass")

	CreateNodeCmd.Flags().String("uuid", "", "Set the nodes uuid")
}
