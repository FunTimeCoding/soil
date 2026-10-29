package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) ResolveChannel(
	identifier string,
	name string,
) (*model.Channel, error) {
	if name != "" {
		return s.TeamChannel(name)
	}

	return s.Channel(identifier)
}
