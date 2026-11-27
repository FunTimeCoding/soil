package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/user"

func (c *Client) AddUser(u *user.User) {
	c.users = append(c.users, u)
}
