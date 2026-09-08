package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Thread(p *model.Post) ([]*post.Post, error) {
	return c.thread[p.Id], nil
}
