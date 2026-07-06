package mattermost

import "github.com/funtimecoding/soil/pkg/chat/mattermost/user_map"

func WithUserMap(m *user_map.Map) Option {
	return func(c *Client) {
		c.user = m
	}
}
