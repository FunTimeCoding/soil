package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/habitica/task"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Tasks(taskType string) ([]*task.Task, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*task.Task, error) {
			return s.client.Tasks(taskType)
		},
	)
}
