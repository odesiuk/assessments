package todolist

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (l todoList) cleanup(_ *cobra.Command, _ []string) error {
	if err := l.repo.Cleanup(); err != nil {
		return fmt.Errorf("cleanup error: %w", err)
	}

	return nil
}
