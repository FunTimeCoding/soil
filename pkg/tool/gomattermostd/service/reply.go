package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) Reply(
	h *model.Channel,
	m *model.Post,
	text string,
) (*model.Post, error) {
	return guarded(
		s.guard,
		func() (*model.Post, error) {
			return s.mattermost.Reply(h, m, text)
		},
	)
}
