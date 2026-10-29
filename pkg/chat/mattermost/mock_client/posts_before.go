package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostsBefore(
	_ *model.Channel,
	_ time.Time,
	_ int,
) ([]*post.Post, error) {
	return nil, c.Failure
}
