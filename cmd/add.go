package cmd

import (
	"fmt"
	"os"

	"github.com/lukastriescoding/open/db"
	"github.com/lukastriescoding/open/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add <path-to-directory> <dir-name>",
	Short: "Add a directory to the db with a name",
	Long:  `The open add command is used to add a directory to the db with a name.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		absolutePath, err := utils.GetAbsolutePath(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = db.InsertDir(args[1], absolutePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Directory added successfully")
	},
}
