package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"strings"
)

func (c *Client) User(identifier string) (*user.User, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, u := range c.users {
		if u.Identifier == identifier ||
			strings.EqualFold(u.Mail, identifier) ||
			strings.EqualFold(u.Principal, identifier) {
			return u, nil
		}
	}

	return nil, not_found.New("user", identifier)
}
