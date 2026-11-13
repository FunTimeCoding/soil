package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) UserThreads(o model.GetUserThreadsOpts) (*model.Threads, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*model.Threads, error) {
			me, e := s.mattermost.Me()

			if e != nil {
				return nil, e
			}

			return s.mattermost.UserThreads(me, s.mattermost.DefaultTeam(), o)
		},
	)
}
