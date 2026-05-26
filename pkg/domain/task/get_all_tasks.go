package task

import "github.com/KeviinMoralees/tasks-api/pkg/domain/entities"

func (s service) GetAllTask() ([]entities.Task, error) {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return []entities.Task{}, err
	}
	return tasks, nil
}
