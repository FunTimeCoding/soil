package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) Me() (*model.User, error) {
	return guarded(s.guard, s.mattermost.Me)
}
