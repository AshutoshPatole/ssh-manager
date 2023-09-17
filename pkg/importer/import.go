/*
Copyright © 2023 AshutoshPatole
*/
package importer

import (
	"fmt"
	"os"

	"github.com/AshutoshPatole/ssh-manager/ssh"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var filePath string

type Group struct {
	Name        string        `yaml:"Name"`
	User        string        `yaml:"User"`
	Environment []Environment `yaml:"Environment"`
}

type Environment struct {
	Name    string   `yaml:"Name"`
	Servers []Server `yaml:"Servers"`
}

type Server struct {
	HostName string `yaml:"HostName"`
}

// importCmd represents the import command
var ImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if filePath != "" {
			readFile(filePath)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")
	ImportCmd.Flags().StringVarP(&filePath, "file", "f", "", "Specify YAML file path to import")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readFile(filepath string) {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}
	var data []Group
	if err := yaml.Unmarshal(yamlFile, &data); err != nil {
		fmt.Printf("Error unmarshaling YAML: %s\n", err)
		return
	}

	for _, group := range data {
		fmt.Println("Importing " + group.Name)
		password := ssh.AskPass()
		for _, env := range group.Environment {
			for _, host := range env.Servers {
				ssh.ConnectServer(host.HostName, group.User, password, group.Name, env.Name)
			}
		}
	}
}