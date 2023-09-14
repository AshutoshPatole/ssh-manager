/*
Copyright © 2023 AshutoshPatole
*/
package cmd

import (
	"fmt"

	"github.com/AshutoshPatole/ssh-manager/pkg/ssh"
	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var aUser, aGroup string

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
		addServer()
	},
}

func init() {
	addCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&aGroup, "group", "g", "default", "Group name in which this server should be added")
	serverCmd.Flags().StringVarP(&aUser, "user", "u", "", "User name to connect")
}

func addServer() {
	password := ssh.AskPass()
	fmt.Println(password)
}
