package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) EditPost(
	p *model.Post,
	message string,
) (*model.Post, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result, okay := c.post[p.Id]

	if !okay {
		return nil, not_found.New(constant.MattermostPostField, p.Id)
	}

	result.Message = message

	return result, nil
}
