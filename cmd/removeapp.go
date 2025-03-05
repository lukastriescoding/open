package cmd

import (
	"fmt"
	"os"

	"git.schonelu.de/lukas/open/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeappCmd)
}

var removeappCmd = &cobra.Command{
	Use:   "remove-app <app-name>",
	Short: "Remove an application with its name",
	Long:  `The remove command is used to remove an application from the db with its name.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := db.RemoveApp(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Application removed successfully")
	},
}
