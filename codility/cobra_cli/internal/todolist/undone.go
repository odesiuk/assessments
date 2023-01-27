package todolist

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (l todoList) undone(_ *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("wrong args provided")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("worng param: %w", err)
	}

	if err := l.repo.Undone(id); err != nil {
		return fmt.Errorf("set undone error: %w", err)
	}

	return nil
}
