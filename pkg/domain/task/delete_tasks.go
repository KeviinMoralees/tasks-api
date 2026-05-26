package task

func (s service) DeleteTask(id int) (string, error) {
	message, err := s.repo.DeleteTask(id)
	if err != nil {
		return "", err
	}
	return message, nil
}
