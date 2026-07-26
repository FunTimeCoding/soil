package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) ChannelUsers(h *model.Channel) ([]*model.User, error) {
	var result []*model.User
	var page int

	for {
		users, _, e := c.client.GetUsersInChannel(
			c.context,
			h.Id,
			page,
			constant.MattermostPerPage,
			constant.MattermostEmptyEntityTag,
		)

		if e != nil {
			return nil, e
		}

		if len(users) == 0 {
			break
		}

		result = append(result, users...)
		page++
	}

	return result, nil
}
