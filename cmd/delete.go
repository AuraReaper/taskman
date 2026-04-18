package cmd

import (
	"fmt"
	"strconv"

	"github.com/AuraReaper/taskman/storage"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(store *storage.Store) *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "del",
		Short: "delete a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteTask(args, store)
		},
	}

	return deleteCmd
}

func deleteTask(args []string, store *storage.Store) error {
	if len(args) == 0 {
		fmt.Println("require task id")
		return nil
	}
	id, _ := strconv.Atoi(args[0])
	return store.Delete(id)
}
