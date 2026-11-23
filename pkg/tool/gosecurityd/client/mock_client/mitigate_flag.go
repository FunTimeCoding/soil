package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/subject"
)

func (c *Client) MitigateFlag(
	_ context.Context,
	name string,
	b client.MitigationRequest,
) (*client.Flag, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	if e := subject.Check(name); e != nil {
		return nil, e
	}

	result := c.findFlag(name)

	if result == nil {
		return nil, not_found.New(constant.FlagKind, name)
	}

	if result.Mitigated {
		return nil, conflict.Format("flag already mitigated: %s", name)
	}

	result.Mitigated = true
	result.Mitigation = b.Mitigation

	return result, nil
}
