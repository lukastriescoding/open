package cmd

import (
	"fmt"
	"os"

	"git.schonelu.de/lukas/open/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listAppsCmd)
}

var listAppsCmd = &cobra.Command{
	Use:   "list-apps",
	Short: "List all the applications",
	Long:  `Print out name and path of all applications.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		apps, err := db.GetAllApps()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var done bool
		for _, app := range apps {
			fmt.Printf("%s: %s\n", app.Name, app.Path)
			done = true
		}
		if !done {
			fmt.Println("No applications found")
		}
	},
}
