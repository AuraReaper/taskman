package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/AuraReaper/taskman/models"
	"github.com/AuraReaper/taskman/storage"
	"github.com/spf13/cobra"
)

func NewAddCmd(store *storage.Store) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			return addTask(args, store)
		},
	}

	return addCmd
}

func addTask(args []string, store *storage.Store) error {
	if len(args) == 0 {
		fmt.Println("please enter the task title")
		return nil
	}
	taskTitle := strings.Join(args, " ")

	id, err := store.GetNextId()
	if err != nil {
		return err
	}

	task := models.Task{
		ID:        id,
		Title:     taskTitle,
		Status:    "PENDING",
		Timestamp: time.Now().UTC(),
	}

	return store.Create(task)
}
