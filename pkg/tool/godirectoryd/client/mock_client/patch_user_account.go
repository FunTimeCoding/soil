package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func (c *Client) PatchUserAccount(
	_ context.Context,
	account string,
	b client.PatchUserAccountJSONRequestBody,
) (*client.User, error) {
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

	u := c.UserList[i]

	if b.Name != nil && *b.Name != "" {
		u.Name = *b.Name
	}

	if b.Surname != nil && *b.Surname != "" {
		u.Surname = *b.Surname
	}

	if b.Mail != nil && *b.Mail != "" {
		u.Mail = *b.Mail
	}

	result := *u

	return &result, nil
}
