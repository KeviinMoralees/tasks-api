package task

import "github.com/KeviinMoralees/tasks-api/pkg/domain/entities"

func (s service) UpdateTask(id int, task entities.Task) (entities.Task, error) {
	task, err := s.repo.UpdateTask(id, task)
	if err != nil {
		return entities.Task{}, err
	}
	return task, nil
}
