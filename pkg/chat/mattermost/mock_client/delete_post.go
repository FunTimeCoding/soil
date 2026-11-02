package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) DeletePost(p *model.Post) error {
	if c.Failure != nil {
		return c.Failure
	}

	if _, okay := c.post[p.Id]; !okay {
		return not_found.New(constant.MattermostPostField, p.Id)
	}

	delete(c.post, p.Id)

	return nil
}
