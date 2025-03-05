package cmd

import (
	"fmt"
	"os"

	"github.com/lukastriescoding/open/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(unsetCmd)
}

var unsetCmd = &cobra.Command{
	Use:   "unset <dir-name>",
	Short: "Unset the main application for a directory",
	Long:  `Unset the main application that this directory will automatically be opened with. It will then be opened with the main-application.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := db.UpdateDirMainApp(args[0], "")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Main application set successfully")
	},
}
