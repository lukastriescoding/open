package cmd

import (
	"fmt"
	"os"

	"github.com/lukastriescoding/open/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mainappCmd)
}

var mainappCmd = &cobra.Command{
	Use:   "main-app <app-name>",
	Short: "Set the main application for open",
	Long:  `Set the main application that any directory will automatically be opened with, if it doesn't have a specific main-app set.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		existsApp, err := db.ExistsApp(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if !existsApp {
			fmt.Println("Application does not exist")
			os.Exit(1)
		}
		err = db.SetMainApp(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Main application set successfully")
	},
}
