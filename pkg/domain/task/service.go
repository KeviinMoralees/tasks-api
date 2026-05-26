package task

import "github.com/KeviinMoralees/tasks-api/pkg/domain/repository"

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{
		repo,
	}
}
