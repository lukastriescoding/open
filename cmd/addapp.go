package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lukastriescoding/open/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addappCmd)
}

var addappCmd = &cobra.Command{
	Use:   "add-app <(path-to-)application> (<app-name>)",
	Short: "Remove a directory with its name",
	Long:  `The remove command is used to remove a directory from the db with its name.`,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			path, err := exec.LookPath(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Executable not found: %v\n", err)
				os.Exit(1)
			}
			err = db.InsertApp("", path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			path, err := exec.LookPath(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Executable not found: %v\n", err)
				os.Exit(1)
			}
			err = db.InsertApp(args[1], path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Println("Application added successfully")
	},
}
