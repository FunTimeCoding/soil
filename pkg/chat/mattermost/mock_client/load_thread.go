package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/thread"
)

func (c *Client) LoadThread(_ *post.Post) *thread.Thread {
	return nil
}
