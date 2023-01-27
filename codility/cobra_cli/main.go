package cobra_cli

import (
	"github.com/odesiuk/assessments/codility/cobra_cli/internal/storage"
	"github.com/odesiuk/assessments/codility/cobra_cli/internal/todolist"
)

func main() {
	if err := todolist.CLI(storage.NewTaskRepository("tasks.json")).Execute(); err != nil {
		panic(err)
	}
}
