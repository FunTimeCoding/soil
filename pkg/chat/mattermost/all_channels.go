package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) AllChannels(
	limit int,
	offset int,
) ([]*model.ChannelWithTeamData, error) {
	if limit <= 0 || limit > constant.MattermostMaxPerPage {
		limit = constant.MattermostMaxPerPage
	}

	if offset < 0 {
		offset = 0
	}

	page := offset / limit
	skip := offset % limit
	result, _, e := c.client.GetAllChannels(
		c.context,
		page,
		limit,
		constant.MattermostEmptyEntityTag,
	)

	if e != nil {
		return nil, wrapError(e)
	}

	if skip == 0 {
		return result, nil
	}

	if skip > len(result) {
		return nil, nil
	}

	result = result[skip:]
	next, _, f := c.client.GetAllChannels(
		c.context,
		page+1,
		limit,
		constant.MattermostEmptyEntityTag,
	)

	if f != nil {
		return nil, wrapError(f)
	}

	result = append(result, next...)

	if len(result) > limit {
		result = result[:limit]
	}

	return result, nil
}
