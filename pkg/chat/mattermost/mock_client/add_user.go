package mock_client

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) AddUser(
	identifier string,
	username string,
) {
	c.user[identifier] = &model.User{Id: identifier, Username: username}
}
