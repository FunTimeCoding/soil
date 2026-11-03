package mock_client

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func (c *Client) PostUser(
	_ context.Context,
	b client.PostUserJSONRequestBody,
) (*client.User, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	if _, e := c.userIndex(b.Account); e == nil {
		return nil, conflict.Exists(constant.UserKind, b.Account)
	}

	result := client.User{
		Account: b.Account,
		Name:    b.Name,
		Surname: b.Surname,
		DistinguishedName: fmt.Sprintf(
			"%s=%s,%s,%s",
			constant.AccountAttribute,
			b.Account,
			constant.UserContainer,
			c.Base,
		),
	}

	if b.Mail != nil {
		result.Mail = *b.Mail
	}

	if b.Password != nil && *b.Password != "" {
		c.Passwords[b.Account] = *b.Password
	}

	c.UserList = append(c.UserList, &result)
	v := result

	return &v, nil
}
