package mock_client

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func (c *Client) PostUserAccountPassword(
	_ context.Context,
	account string,
	b client.PostUserAccountPasswordJSONRequestBody,
) (*client.Receipt, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	if _, e := c.userIndex(account); e != nil {
		return nil, e
	}

	c.Passwords[account] = b.Password

	return &client.Receipt{Message: fmt.Sprintf("password set: %s", account)}, nil
}
