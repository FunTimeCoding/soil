package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
)

func (c *Client) userIndex(account string) (int, error) {
	for i, u := range c.UserList {
		if u.Account == account {
			return i, nil
		}
	}

	return 0, not_found.New(constant.UserKind, account)
}
