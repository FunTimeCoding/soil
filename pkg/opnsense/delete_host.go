package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (c *Client) DeleteHost(identifier string) error {
	return postDelete(
		c,
		constant.HostSubject,
		join.Slash([]string{constant.HostDelete, identifier}),
		identifier,
	)
}
