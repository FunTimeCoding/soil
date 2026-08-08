package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.MattermostHostEnvironment); s != "" {
		o = append(o, WithHost(s))
	}

	if s := environment.Optional(constant.MattermostTokenEnvironment); s != "" {
		o = append(o, WithToken(s))
	}

	if s := environment.Optional(constant.MattermostTeamEnvironment); s != "" {
		o = append(o, WithTeam(s))
	}

	if s := environment.Optional(constant.MattermostChannelEnvironment); s != "" {
		o = append(o, WithChannel(s))
	}

	if environment.Exists(constant.MattermostInsecureEnvironment) {
		o = append(o, WithInsecure())
	}

	return New(o...)
}
