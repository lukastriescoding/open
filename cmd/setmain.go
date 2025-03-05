package cmd

import (
	"fmt"
	"os"

	"git.schonelu.de/lukas/open/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setmainCmd)
}

var setmainCmd = &cobra.Command{
	Use:   "set-main <dir-name> <app-name>",
	Short: "Set the main application for a directory",
	Long:  `Set the main application that this directory will automatically be opened with.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		existsApp, err := db.ExistsApp(args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if !existsApp {
			fmt.Println("Application does not exist")
			os.Exit(1)
		}
		err = db.UpdateDirMainApp(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Main application set successfully")
	},
}
