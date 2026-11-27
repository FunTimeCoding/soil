package mock_client

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
	"strings"
)

func (c *Client) Users(
	query string,
	limit int,
) ([]*user.User, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	q := strings.ToLower(query)
	var result []*user.User

	for _, u := range c.users {
		if strings.HasPrefix(strings.ToLower(u.Name), q) ||
			strings.HasPrefix(strings.ToLower(u.Mail), q) ||
			strings.HasPrefix(strings.ToLower(u.Principal), q) {
			result = append(result, u)
		}
	}

	return limited(result, limit), nil
}
