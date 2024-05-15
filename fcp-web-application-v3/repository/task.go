package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
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
	// return nil // TODO: replace this
	return t.filebased.UpdateTask(task.ID, *task)
}

func (t *taskRepository) Delete(id int) error {
	// return nil // TODO: replace this
	err := t.filebased.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	// return nil, nil // TODO: replace this
	task, err := t.filebased.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	// return nil, nil // TODO: replace this
	// return t.filebased.GetTasks()
	task, err := t.filebased.GetTasks()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	// return nil, nil // TODO: replace this
	task, err := t.filebased.GetTaskListByCategory(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}
