/*
Copyright © 2023 AshutoshPatole
*/
package list

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list groups and servers",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("list called")
	// },
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func ListGroups() {
// 	var config c.Config

// 	if err := viper.Unmarshal(&config); err != nil {
// 		log.Fatalln(err)
// 	}

// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader([]string{"ID.", "Group Name", "Server(s)"})
// 	servers := 0
// 	for idx, group := range config.Groups {
// 		servers = len(group.Servers)
// 		table.Append([]string{fmt.Sprint(idx), group.Name, fmt.Sprint(servers)})
// 	}
// 	table.Render()
// }

// func ListServers(group string) {
// 	var config c.Config

// 	if err := viper.Unmarshal(&config); err != nil {
// 		log.Fatalln(err)
// 	}
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader([]string{"ID", "Server", "IP", "Key status"})

// 	// ANSI escape code for colors
// 	const (
// 		green = "\033[32m"
// 		red   = "\033[31m"
// 		reset = "\033[0m"
// 	)

// 	// Checkmark and cross characters
// 	const (
// 		checkmark = "✓"
// 		cross     = "✗"
// 	)

// 	for _, grp := range config.Groups {
// 		if grp.Name == group {
// 			for idx, server := range grp.Servers {
// 				status := ""
// 				if server.KeyAuth {
// 					status = green + checkmark + reset
// 				} else {
// 					status = red + cross + reset
// 				}
// 				table.Append([]string{fmt.Sprint(idx + 1), server.HostName, server.IP, status})
// 			}
// 		}
// 	}
// 	table.Render()
// }
