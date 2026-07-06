package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.HostEnvironment); s != "" {
		o = append(o, WithHost(s))
	}

	if s := environment.Optional(constant.TokenEnvironment); s != "" {
		o = append(o, WithToken(s))
	}

	if s := environment.Optional(constant.TeamEnvironment); s != "" {
		o = append(o, WithTeam(s))
	}

	if s := environment.Optional(constant.ChannelEnvironment); s != "" {
		o = append(o, WithChannel(s))
	}

	return New(o...)
}
