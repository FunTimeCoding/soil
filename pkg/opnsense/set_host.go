package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (c *Client) SetHost(
	identifier string,
	h *request.Host,
) error {
	_, e := postSave(
		c,
		constant.HostSubject,
		join.Slash([]string{constant.HostSet, identifier}),
		request.NewHostWrapper(h),
	)

	return e
}
