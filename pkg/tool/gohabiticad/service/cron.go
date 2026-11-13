package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/habitica/cron"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Cron() (*cron.Cron, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*cron.Cron, error) {
			return s.client.Cron()
		},
	)
}
