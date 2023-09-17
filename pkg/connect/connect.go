/*
Copyright © 2023 AshutoshPatole
*/
package connect

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AshutoshPatole/ssh-manager/ssh"
	c "github.com/AshutoshPatole/ssh-manager/utils"
	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cGroup string

// connectCmd represents the connect command
var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ListToConnectServers(cGroup)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	ConnectCmd.Flags().StringVarP(&cGroup, "group", "g", "", "Specify which group servers should be displayed")
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ListToConnectServers(group string) {
	var config c.Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}

	var servers []string
	var user string
	var environment string

	for _, grp := range config.Groups {
		if grp.Name == group {
			user = grp.User
			for _, env := range grp.Environment {
				environment = env.Name
				for _, server := range env.Servers {
					servers = append(servers, server.HostName)
				}
			}
		}
	}
	toConnect := ""
	prompt := &survey.Select{
		Message: "Select server",
		Options: servers,
	}
	survey.AskOne(prompt, &toConnect)
	fmt.Println(color.InGreen("Trying to connect to " + toConnect))
	ssh.Connect(toConnect, user, environment)

}
