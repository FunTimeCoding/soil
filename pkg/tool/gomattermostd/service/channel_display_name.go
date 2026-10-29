package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (s *Service) ChannelDisplayName(c *model.Channel) string {
	if c.DisplayName != "" {
		return c.DisplayName
	}

	if c.Type != model.ChannelTypeDirect {
		return c.Name
	}

	me, e := s.Me()

	if e != nil {
		return c.Name
	}

	for _, p := range strings.Split(c.Name, "__") {
		if p != me.Id {
			other, f := s.User(p)

			if f != nil || other == nil {
				return c.Name
			}

			return other.Username
		}
	}

	return join.Empty(me.Username, " (self)")
}
