/*
Copyright © 2025 LUKAS <me@schonelu.de>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lukastriescoding/open/db"
	"github.com/lukastriescoding/open/models"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "open (<app-name>) <dir-name>",
	Short: "simplify opening directories with applications",
	Long:  `open is a tool to easily open any directories with an application of your choice.`,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var app models.Application
		var dir models.Directory
		if len(args) == 1 {
			var err error
			dir, err = db.GetDir(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			appSet := false
			if dir.MainApp.Valid && dir.MainApp.String != "" {
				app, err = db.GetApp(dir.MainApp.String)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				appSet = true
			}
			if !appSet {
				app, err = db.GetMainApp()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				appSet = true
			}
			if !appSet {
				fmt.Println("No main application for open or directory set.")
				os.Exit(1)
			}
		} else {
			var err error
			app, err = db.GetApp(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			dir, err = db.GetDir(args[1])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		command := exec.Command(app.Path, dir.Path)
		err := command.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.open.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//rootCmd.AddCommand(addCmd)
}
