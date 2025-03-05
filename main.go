/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"git.schonelu.de/lukas/open/cmd"
	"git.schonelu.de/lukas/open/db"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := db.InitCon()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	cmd.Execute()
}
