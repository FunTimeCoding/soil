package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewEnvironment(p ...Option) *Client {
	l := locator.New(environment.Required(constant.HostEnvironment))

	if v := environment.FallbackInteger(constant.PortEnvironment, 0); v != 0 {
		l.Port(v)
	}

	if environment.Exists(constant.InsecureEnvironment) {
		l.Insecure()
	}

	if h := environment.Optional(constant.PublicHostEnvironment); h != "" {
		p = append([]Option{WithPublicHost(h)}, p...)
	}

	return New(l, environment.Required(constant.TokenEnvironment), p...)
}
