package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) AddReply(
	root string,
	identifier string,
) {
	c.thread[root] = append(
		c.thread[root],
		post.New(&model.Post{Id: identifier, RootId: root}),
	)
}
