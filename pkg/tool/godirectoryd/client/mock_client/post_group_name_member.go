package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"slices"
)

func (c *Client) PostGroupNameMember(
	_ context.Context,
	name string,
	b client.PostGroupNameMemberJSONRequestBody,
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

	if _, f := c.userIndex(b.Account); f != nil {
		return nil, f
	}

	g := c.GroupList[i]

	if slices.Contains(g.Member, b.Account) {
		return nil, conflict.Exists(constant.MemberAttribute, b.Account)
	}

	g.Member = append(g.Member, b.Account)

	return groupCopy(g), nil
}
