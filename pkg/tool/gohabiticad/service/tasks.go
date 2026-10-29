package service

import "github.com/funtimecoding/soil/pkg/habitica/task"

func (s *Service) Tasks(taskType string) ([]*task.Task, error) {
	return guarded(
		s.guard,
		func() ([]*task.Task, error) {
			return s.client.Tasks(taskType)
		},
	)
}
