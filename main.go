/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/AuraReaper/taskman/cmd"
	"github.com/AuraReaper/taskman/storage"
)

func main() {
	store, err := storage.NewStore("data.json")
	if err != nil {
		log.Fatal(err)
	}

	rootCmd := cmd.NewRootCmd(store)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
