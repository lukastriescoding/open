package cmd

import (
	"fmt"
	"os"

	"github.com/lukastriescoding/open/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove <dir-name>",
	Short: "Remove a directory with its name",
	Long:  `The remove command is used to remove a directory from the db with its name.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := db.RemoveDir(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Directory removed successfully")
	},
}
