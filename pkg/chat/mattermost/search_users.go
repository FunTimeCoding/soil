package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) SearchUsers(
	t *model.Team,
	query string,
	limit int,
) ([]*model.User, error) {
	result, r, e := c.client.AutocompleteUsersInTeam(
		c.context,
		t.Id,
		query,
		limit,
		constant.MattermostEmptyEntityTag,
	)

	if e != nil {
		return nil, wrapError(r, e)
	}

	return result.Users, nil
}
