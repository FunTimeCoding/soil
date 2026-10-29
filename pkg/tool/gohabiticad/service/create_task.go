package service

import (
	"github.com/funtimecoding/soil/pkg/habitica/request"
	"github.com/funtimecoding/soil/pkg/habitica/task"
)

func (s *Service) CreateTask(b *request.CreateTaskBody) (*task.Task, error) {
	return guarded(
		s.guard,
		func() (*task.Task, error) {
			return s.client.CreateTask(b)
		},
	)
}
