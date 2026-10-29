package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) React(
	p *model.Post,
	emoji string,
) error {
	return attempt(
		s.guard,
		func() error {
			return s.mattermost.React(p, emoji)
		},
	)
}
