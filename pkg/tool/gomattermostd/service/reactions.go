package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) Reactions(p *model.Post) ([]*model.Reaction, error) {
	return guarded(
		s.guard,
		func() ([]*model.Reaction, error) {
			return s.mattermost.Reactions(p)
		},
	)
}
