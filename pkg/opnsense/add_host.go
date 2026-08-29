package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
)

func (c *Client) AddHost(h *request.Host) (string, error) {
	result, e := postSave(
		c,
		constant.HostSubject,
		constant.HostAdd,
		request.NewHostWrapper(h),
	)

	if e != nil {
		return "", e
	}

	return result.Identifier, nil
}
