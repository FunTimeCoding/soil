package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) AddPost(identifier string) {
	c.post[identifier] = &model.Post{Id: identifier}
}
