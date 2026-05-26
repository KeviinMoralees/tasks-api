package task

func (s service) CompleteTask(id int) (string, error) {
	message, err := s.repo.CompleteTask(id)
	if err != nil {
		return "", nil
	}
	return message, nil
}
