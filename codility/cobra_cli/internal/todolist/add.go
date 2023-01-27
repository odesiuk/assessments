package todolist

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/odesiuk/assessments/codility/cobra_cli/internal/models"
)

func (l todoList) add(_ *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("wrong args provided")
	}

	if err := l.repo.Add(models.Task{Name: args[0]}); err != nil {
		return fmt.Errorf("add task error: %w", err)
	}

	return nil
}
