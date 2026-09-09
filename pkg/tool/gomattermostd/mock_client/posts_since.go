package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostsSince(
	_ *model.Channel,
	_ time.Time,
) ([]*post.Post, error) {
	return nil, nil
}
