package task

import (
	"github.com/KeviinMoralees/tasks-api/pkg/domain/entities"
	"github.com/KeviinMoralees/tasks-api/pkg/domain/repository"
)

type Service interface {
	GetAllTasks() ([]entities.Task, error)
	CreateTask(entities.Task) (entities.Task, error)
	UpdateTask(id int, task entities.Task) (entities.Task, error)
	DeleteTask(id int) (string, error)
	CompleteTask(id int) (string, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return service{
		repo,
	}
}
