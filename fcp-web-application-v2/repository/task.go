package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"

	"go.etcd.io/bbolt"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(taskID int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	t.filebased.StoreTask(*task)

	return nil
}

func (t *taskRepository) Update(taskID int, task *model.Task) error {
	taskJSON, err := json.Marshal(task)
	if err != nil {
		return err
	}
	err = t.filebased.UpdateTask(taskID, task)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) Delete(id int) error {
	taskJSON, err := json.Marshal([]int{id})
	if err != nil {
		return err
	}
	err = t.filebased.DeleteTask("tasks", taskJSON)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	task, err := t.filebased.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	var tasks []model.Task
	err := t.filebased.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if b == nil {
			return fmt.Errorf("tasks bucket not found")
		}
		return b.ForEach(func(k, v []byte) error {
			var task model.Task
			if err := json.Unmarshal(v, &task); err != nil {
				return fmt.Errorf("error unmarshalling task: %v", err)
			}
			tasks = append(tasks, task)
			return nil
		})
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks: %v", err)
	}
	return tasks, nil
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	var taskCategories []model.TaskCategory
	err := t.filebased.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if b == nil {
			return fmt.Errorf("tasks bucket not found")
		}
		return b.ForEach(func(k, v []byte) error {
			var task model.Task
			if err := json.Unmarshal(v, &task); err != nil {
				return fmt.Errorf("error unmarshaling task: %v", err)
			}
			taskCategories = append(taskCategories, model.TaskCategory{
				ID:       task.ID,
				Title:    task.Title,
				Category: task.Category,
			})
			return nil
		})
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks for category %d: %v", id, err)
	}
	if len(taskCategories) == 0 {
		return nil, fmt.Errorf("no tasks found for category ID: %d", id)
	}
	return taskCategories, nil
}
