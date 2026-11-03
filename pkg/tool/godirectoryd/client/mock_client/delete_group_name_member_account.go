package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"slices"
)

func (c *Client) DeleteGroupNameMemberAccount(
	_ context.Context,
	name string,
	account string,
) (*client.Group, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	i, e := c.groupIndex(name)

	if e != nil {
		return nil, e
	}

	g := c.GroupList[i]

	if !slices.Contains(g.Member, account) {
		return nil, not_found.New(constant.MemberAttribute, account)
	}

	g.Member = slices.DeleteFunc(
		g.Member,
		func(m string) bool { return m == account },
	)

	return groupCopy(g), nil
}
