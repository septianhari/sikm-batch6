package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"encoding/json"
	"fmt"

	"go.etcd.io/bbolt"
)

type TaskService interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskService struct {
	taskRepository repo.TaskRepository
}

func NewTaskService(taskRepository repo.TaskRepository) TaskService {
	return &taskService{taskRepository}
}

func (c *taskService) Store(task *model.Task) error {
	err := c.taskRepository.Store(task)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) Update(id int, task *model.Task) error {
	updatedTask, err := s.taskRepository.GetByID(id)
	if err != nil {
		return err
	}

	updatedTask.Title = task.Title
	updatedTask.Deadline = task.Deadline
	updatedTask.Priority = task.Priority
	updatedTask.CategoryID = task.CategoryID
	updatedTask.Status = task.Status

	err = s.taskRepository.Update(id, updatedTask)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) Delete(id int) error {
	err := s.taskRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) GetByID(id int) (*model.Task, error) {
	task, err := s.taskRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetList() ([]model.Task, error) {
	tasks, err := s.taskRepository.GetList()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *taskService) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	var taskCategories []model.TaskCategory
	err := s.taskRepository.DB.View(func(tx *bbolt.Tx) error {
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
