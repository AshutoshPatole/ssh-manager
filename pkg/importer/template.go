/*
Copyright © 2023 AshutoshPatole
*/
package importer

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "save yaml template for bulk import of servers",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("template called")
		saveTemplate()
	},
}

func init() {
	ImportCmd.AddCommand(templateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// templateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func saveTemplate() {
	content := `
- name: groupname
  user: username
  env:
    - name: dev|uat|sit|ppd|prd
      servers:
        - hostname: gvx0lcmdme01d.aholdusa.com
          alias: dev engine
        - hostname: gvx0lcmdms01d.aholdusa.com
          alias: dev service
        - hostname: gvx0lcmdme01p.aholdusa.com
          alias: prod engine

- name: groupname2
  user: anotheruser
  env:
    - name: dev|uat|sit|ppd|prd
      servers:
        - hostname: gvx0lcmdme01d.aholdusa.com
          alias: dev engine
        - hostname: gvx0lcmdms01d.aholdusa.com
          alias: dev service
        - hostname: gvx0lcmdme01p.aholdusa.com
          alias: prod engine
`
	// fmt.Println(content)
	data := []Group{}
	err := yaml.Unmarshal([]byte(content), &data)
	if err != nil {
		fmt.Println("something wrong in unmarshaling", err.Error())
		return
	}
	d, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Println("something wrong in marshaling")
		return
	}
	fmt.Println(string(d))
	homeDir, _ := os.UserHomeDir()
	wErr := os.WriteFile(homeDir+"/ssh-import-template.yaml", d, 0666)
	if wErr != nil {
		fmt.Println("error while saving file", wErr.Error())
		return
	}
	fmt.Println(color.InGreen("File saved at " + homeDir + "/ssh-import-template.yaml"))
}
