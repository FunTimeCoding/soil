package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func (c *Client) GetUser(_ context.Context) ([]*client.User, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	result := []*client.User{}

	for _, u := range c.UserList {
		v := *u
		result = append(result, &v)
	}

	return result, nil
}
