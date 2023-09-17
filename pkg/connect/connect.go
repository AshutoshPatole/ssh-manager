/*
Copyright © 2023 AshutoshPatole
*/
package connect

import (
	"fmt"
	"log"
	"strings"

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

type ServerOption struct {
	Label       string
	Environment string
	HostName    string
}

func ListToConnectServers(group string) {
	var config c.Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}

	selectedEnvName := ""
	selectedHostName := ""

	serverOptions := []ServerOption{}
	user := ""

	for _, grp := range config.Groups {
		if grp.Name == group {
			user = grp.User
			for _, env := range grp.Environment {
				for _, server := range env.Servers {
					serverOption := ServerOption{
						Label:       fmt.Sprintf("%s (%s)", server.HostName, env.Name),
						Environment: env.Name,
						HostName:    server.HostName,
					}
					serverOptions = append(serverOptions, serverOption)
				}
			}
		}
	}
	labels := make([]string, len(serverOptions))
	for i, serverOption := range serverOptions {
		labels[i] = serverOption.Label
	}

	fmt.Println(color.InGreen(selectedEnvName))
	prompt := &survey.Select{
		Message: "Select server",
		Options: labels,
	}
	survey.AskOne(prompt, &selectedHostName)

	// Extract environment name from the selected option
	for _, serverOption := range serverOptions {
		if serverOption.Label == selectedHostName {
			selectedEnvName = serverOption.Environment
			selectedHostName = strings.Split(serverOption.HostName, " (")[0]
			break
		}
	}

	fmt.Println(color.InGreen("Trying to connect to " + selectedHostName))
	ssh.Connect(selectedHostName, user, selectedEnvName)

}
