/*
Copyright © 2023 AshutoshPatole
*/
package add

import (
	"fmt"
	"os"

	"github.com/AshutoshPatole/ssh-manager/ssh"
	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var aUser, aGroup, aEnv string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use: "server",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf(color.InRedOverBlack("Expects server name or ip address"))
		}
		return nil
	},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		isValid := false
		for _, env := range allowedEnvironments {
			if aEnv == env {
				isValid = true
			}
		}
		if isValid {
			addServer(args[0], aUser)
		} else {
			fmt.Print(color.InYellow("Unknown environment: Allowed values are "))
			fmt.Println(allowedEnvironments)
			os.Exit(1)
		}
	},
}

func init() {
	AddCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&aGroup, "group", "g", "", "Group name in which this server should be added")
	serverCmd.MarkFlagRequired("group")
	serverCmd.Flags().StringVarP(&aUser, "user", "u", "", "User name to connect")
	serverCmd.MarkFlagRequired("user")
	serverCmd.Flags().StringVarP(&aEnv, "env", "e", "dev", "Environment to add ")
	serverCmd.MarkFlagRequired("env")
}

func addServer(server, user string) {
	password := ssh.AskPass()
	ssh.InitServer(server, user, password, aGroup, aEnv)
}
