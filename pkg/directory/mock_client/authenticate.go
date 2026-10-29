package mock_client

import (
	"github.com/funtimecoding/soil/pkg/directory/constant"
	"github.com/funtimecoding/soil/pkg/directory/types/entry"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (c *Client) Authenticate(
	account string,
	password string,
) (*entry.Entry, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, n := range c.entries {
		if n.Account != account && n.Mail != account {
			continue
		}

		if c.passwords[n.Account] != password {
			return nil, validation.New("invalid credentials")
		}

		return n, nil
	}

	return nil, not_found.New(constant.Subject, account)
}
