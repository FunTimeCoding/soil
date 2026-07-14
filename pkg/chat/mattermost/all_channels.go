package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) AllChannels(
	limit int,
	offset int,
) ([]*model.ChannelWithTeamData, error) {
	if limit <= 0 || limit > constant.MaxPerPage {
		limit = constant.MaxPerPage
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
		constant.EmptyEntityTag,
	)

	if e != nil {
		return nil, e
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
		constant.EmptyEntityTag,
	)

	if f != nil {
		return nil, f
	}

	result = append(result, next...)

	if len(result) > limit {
		result = result[:limit]
	}

	return result, nil
}
