package cmd

import (
	"fmt"
	"os"

	"git.schonelu.de/lukas/open/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the directories",
	Long:  `Print out name, path and main-app of all directories.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		dirs, err := db.GetAllDirs()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var done bool
		for _, dir := range dirs {
			fmt.Printf("%s: %s\n", dir.Name, dir.Path)
			if dir.MainApp.Valid && dir.MainApp.String != "" {
				fmt.Printf("  Main app: %s\n", dir.MainApp.String)
			}
			done = true
		}
		if !done {
			fmt.Println("No directories found")
		}
	},
}
