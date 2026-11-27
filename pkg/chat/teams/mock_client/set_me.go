package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/user"

func (c *Client) SetMe(u *user.User) {
	c.me = u
}
