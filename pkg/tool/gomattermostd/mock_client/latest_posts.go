package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) LatestPosts(
	_ *model.Channel,
	_ int,
) ([]*post.Post, error) {
	return nil, nil
}
