package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) DirectMessage(
	u *model.User,
	text string,
) (*model.Post, error) {
	return guarded(
		s.guard,
		func() (*model.Post, error) {
			return s.mattermost.DirectMessage(u, text)
		},
	)
}
