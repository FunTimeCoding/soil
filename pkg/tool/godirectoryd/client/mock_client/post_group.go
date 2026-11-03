package mock_client

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
)

func (c *Client) PostGroup(
	_ context.Context,
	b client.PostGroupJSONRequestBody,
) (*client.Group, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	if _, e := c.groupIndex(b.Name); e == nil {
		return nil, conflict.Exists(constant.GroupKind, b.Name)
	}

	number := constant.FirstGroupNumber

	for _, g := range c.GroupList {
		if g.Number >= number {
			number = g.Number + 1
		}
	}

	result := &client.Group{
		Name:   b.Name,
		Number: number,
		Member: []string{},
		DistinguishedName: fmt.Sprintf(
			"%s=%s,%s,%s",
			constant.NameAttribute,
			b.Name,
			constant.GroupContainer,
			c.Base,
		),
	}
	c.GroupList = append(c.GroupList, result)

	return groupCopy(result), nil
}
