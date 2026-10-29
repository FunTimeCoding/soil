package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) PostBeforeMilli(
	_ *model.Channel,
	_ int64,
) (*post.Post, error) {
	return nil, c.Failure
}
