package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) FindPostBefore(
	_ *model.Channel,
	_ time.Time,
) (*post.Post, bool, error) {
	return nil, false, c.Failure
}
