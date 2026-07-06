package mattermost

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/user_map"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(o ...Option) *Client {
	result := &Client{context: context.Background()}

	for _, p := range o {
		p(result)
	}

	if result.user == nil {
		result.user = user_map.New(nil, []string{})
	}

	if result.host == "" {
		panic("host required")
	}

	if result.token == "" {
		panic("token required")
	}

	result.client = model.NewAPIv4Client(
		locator.New(result.host).String(),
	)
	result.client.SetOAuthToken(result.token)

	if result.teamName != "" {
		result.team = result.MustTeam(result.teamName)

		if result.channelName != "" {
			result.channel = result.MustChannelByName(
				result.team,
				result.channelName,
			)
		}
	}

	result.RefreshSocket()

	return result
}
