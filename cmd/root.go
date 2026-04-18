package cmd

import (
	"github.com/AuraReaper/taskman/storage"
	"github.com/spf13/cobra"
)

func NewRootCmd(store *storage.Store) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "taskman",
		Short: "Task manager CLI",
	}

	rootCmd.AddCommand(NewAddCmd(store))
	rootCmd.AddCommand(NewListCmd(store))
	rootCmd.AddCommand(NewCompleteCmd(store))
	rootCmd.AddCommand(NewDeleteCmd(store))

	return rootCmd
}
