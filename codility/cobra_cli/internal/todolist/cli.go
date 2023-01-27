package todolist

import (
	"github.com/spf13/cobra"

	"github.com/odesiuk/assessments/codility/cobra_cli/internal/models"
)

const name = "todolist"

type TaskRepository interface {
	Add(task models.Task) error
	List() ([]models.Task, error)
	Done(id int) error
	Undone(id int) error
	Cleanup() error
}

type todoList struct {
	repo TaskRepository
}

func CLI(repo TaskRepository) *cobra.Command {
	rootCmd := &cobra.Command{Use: name}

	tdl := todoList{repo}

	// init commands
	rootCmd.AddCommand(&cobra.Command{
		Use:   "add [Task name]",
		Short: "Add task to the list",
		Args:  cobra.MinimumNArgs(1),
		RunE:  tdl.add,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List all tasks still to do",
		Args:  cobra.NoArgs,
		RunE:  tdl.list,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "done [task id]",
		Short: "Mark task as done",
		Args:  cobra.MinimumNArgs(1),
		RunE:  tdl.done,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "undone [task id]",
		Short: "Mark task as not done",
		Args:  cobra.MinimumNArgs(1),
		RunE:  tdl.undone,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "cleanup",
		Short: "Cleanup done tasks",
		Args:  cobra.NoArgs,
		RunE:  tdl.cleanup,
	})

	return rootCmd
}
