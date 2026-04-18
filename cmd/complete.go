package cmd

import (
	"fmt"
	"strconv"

	"github.com/AuraReaper/taskman/storage"
	"github.com/spf13/cobra"
)

func NewCompleteCmd(store *storage.Store) *cobra.Command {
	completeCmd := &cobra.Command{
		Use:   "complete",
		Short: "complete a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return completeTask(args, store)
		},
	}

	return completeCmd
}

func completeTask(args []string, store *storage.Store) error {
	if len(args) == 0 {
		fmt.Println("require task id")
		return nil
	}
	id, _ := strconv.Atoi(args[0])
	return store.UpdateStatus(id)
}
