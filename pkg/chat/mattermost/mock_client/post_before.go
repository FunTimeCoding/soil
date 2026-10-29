package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostBefore(
	_ *model.Channel,
	_ time.Time,
) (*post.Post, error) {
	return nil, c.Failure
}
