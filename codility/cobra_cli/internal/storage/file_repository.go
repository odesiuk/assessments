package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/odesiuk/assessments/codility/cobra_cli/internal/models"
)

type TaskRepository struct {
	path string
}

func NewTaskRepository(path string) *TaskRepository {
	r := &TaskRepository{path: path}

	if _, err := r.List(); err != nil {
		_ = r.save(make([]models.Task, 0))
	}

	return r
}

func (r TaskRepository) Add(task models.Task) error {
	list, err := r.List()
	if err != nil {
		return err
	}

	return r.save(append(list, task))
}

func (r TaskRepository) save(list []models.Task) error {
	content, err := json.Marshal(list)
	if err != nil {
		return fmt.Errorf("marshall error: %w", err)
	}

	err = ioutil.WriteFile(r.path, content, 0644)
	if err != nil {
		return fmt.Errorf("write to file error: %w", err)
	}

	return nil
}

func (r TaskRepository) List() ([]models.Task, error) {
	content, err := ioutil.ReadFile(r.path)
	if err != nil {
		return nil, fmt.Errorf("read from file error: %w", err)
	}

	var list []models.Task

	if err = json.Unmarshal(content, &list); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return list, nil
}

func (r TaskRepository) Done(id int) error {
	list, err := r.List()
	if err != nil {
		return err
	}

	if id > 0 && len(list) >= id {
		list[id-1].Done = true
	}

	return r.save(list)
}

func (r TaskRepository) Undone(id int) error {
	list, err := r.List()
	if err != nil {
		return err
	}

	if id > 0 && len(list) >= id {
		list[id-1].Done = false
	}

	return r.save(list)
}

func (r TaskRepository) Cleanup() error {
	list, err := r.List()
	if err != nil {
		return err
	}

	cleaned := make([]models.Task, 0)

	for _, task := range list {
		if !task.Done {
			cleaned = append(cleaned, task)
		}
	}

	return r.save(cleaned)
}
