package mock_client

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"slices"
)

func (c *Client) DeleteUserAccount(
	_ context.Context,
	account string,
) (*client.Receipt, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	i, e := c.userIndex(account)

	if e != nil {
		return nil, e
	}

	c.UserList = slices.Delete(c.UserList, i, i+1)
	delete(c.Passwords, account)

	return &client.Receipt{Message: fmt.Sprintf("deleted user: %s", account)}, nil
}
