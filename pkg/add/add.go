/*
Copyright © 2023 AshutoshPatole
*/
package add

import (
	"github.com/spf13/cobra"
)

var allowedEnvironments = []string{"dev", "uat", "sit", "ppd", "prd"}

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add servers and groups in configurations",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
