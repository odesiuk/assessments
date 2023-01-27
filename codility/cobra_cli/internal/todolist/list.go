package todolist

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (l todoList) list(_ *cobra.Command, _ []string) error {
	list, err := l.repo.List()
	if err != nil {
		return fmt.Errorf("get list error: %w", err)
	}

	for i, task := range list {
		if !task.Done {
			fmt.Println(fmt.Sprintf("%d: %s", i+1, task.Name))
		}
	}

	return nil
}
