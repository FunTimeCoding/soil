package request_context

import (
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (c *Context) Address() string {
	if v, okay := c.Header()[constant.RealAddress]; okay {
		return v[0]
	}

	return network.SplitHost(c.request.RemoteAddr)
}
