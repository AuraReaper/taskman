package cmd

import (
	"fmt"

	"github.com/AuraReaper/taskman/storage"
	"github.com/spf13/cobra"
)

func NewListCmd(store *storage.Store) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "list all tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listTasks(store)
		},
	}

	return listCmd
}

func listTasks(store *storage.Store) error {
	tasks, err := store.ReadAll()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		fmt.Printf("ID:- %v\nTask Name:-  %s \nStatus:- %s\n", task.ID, task.Title, task.Status)
	}

	return nil
}
