package service

import "github.com/funtimecoding/soil/pkg/chat/teams/channel"

func (s *Service) Channels(team string) ([]*channel.Channel, error) {
	identifier, e := s.resolveTeam(team)

	if e != nil {
		return nil, e
	}

	return s.teamChannels(identifier)
}
