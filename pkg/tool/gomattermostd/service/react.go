package service

import (
	"context"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) React(
	p *model.Post,
	emoji string,
) error {
	return s.guard.Attempt(
		context.Background(),
		func(context.Context) error {
			return s.mattermost.React(p, emoji)
		},
	)
}
