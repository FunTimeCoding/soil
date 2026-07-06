package model_context

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (s *Server) channelDisplayName(c *model.Channel) string {
	if c.DisplayName != "" {
		return c.DisplayName
	}

	if c.Type == model.ChannelTypeDirect {
		me, e := s.client.Me()

		if e != nil {
			return c.Name
		}

		parts := strings.Split(c.Name, "__")

		for _, p := range parts {
			if p != me.Id {
				other, f := s.client.User(p)

				if f != nil || other == nil {
					return c.Name
				}

				return other.Username
			}
		}

		return join.Empty(me.Username, " (self)")
	}

	return c.Name
}
