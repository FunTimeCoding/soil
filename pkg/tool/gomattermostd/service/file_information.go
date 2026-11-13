package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) FileInformation(identifier string) (*model.FileInfo, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*model.FileInfo, error) {
			return s.mattermost.FileInformation(identifier)
		},
	)
}
