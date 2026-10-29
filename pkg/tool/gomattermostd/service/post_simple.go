package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) PostSimple(
	h *model.Channel,
	text string,
) (*model.Post, error) {
	return guarded(
		s.guard,
		func() (*model.Post, error) {
			return s.mattermost.PostSimple(h, text)
		},
	)
}
