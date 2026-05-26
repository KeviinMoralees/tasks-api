package task

import "github.com/KeviinMoralees/tasks-api/pkg/domain/entities"

func (s *Service) CreateTask(task entities.Task) (entities.Task, error) {
	task, err := s.CreateTask(task)
	if err != nil {
		return entities.Task{}, err
	}

	return task, nil
}
