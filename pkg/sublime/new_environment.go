package sublime

import (
	"github.com/funtimecoding/soil/pkg/sublime/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.HostEnvironment); s != "" {
		o = append(o, WithHost(s))
	}

	return New(o...)
}
