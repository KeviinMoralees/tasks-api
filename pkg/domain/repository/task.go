package repository

import "github.com/KeviinMoralees/tasks-api/pkg/domain/entities"

type Repository interface {
	GetAllTasks() ([]entities.Task, error)
	CreateTask(entities.Task) (entities.Task, error)
	UpdateTask(id int, task entities.Task) (entities.Task, error)
	DeleteTask(id int) (string, error)
	CompleteTask(id int) (string, error)
}
