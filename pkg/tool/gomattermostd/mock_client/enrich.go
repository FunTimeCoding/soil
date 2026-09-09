package mock_client

import "github.com/funtimecoding/soil/pkg/chat/mattermost/post"

func (c *Client) Enrich(_ []*post.Post) error {
	return nil
}
