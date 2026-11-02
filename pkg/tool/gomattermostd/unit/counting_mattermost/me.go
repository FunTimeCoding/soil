package counting_mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Me() (*model.User, error) {
	c.me.Add(1)

	return c.Client.Me()
}
