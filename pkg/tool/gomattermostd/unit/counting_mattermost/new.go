package counting_mattermost

import "github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"

func New(c *mock_client.Client) *Client {
	return &Client{Client: c}
}
