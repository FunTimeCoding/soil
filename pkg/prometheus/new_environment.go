package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.FallbackInteger(constant.PortEnvironment, 0),
		!environment.Exists(constant.InsecureEnvironment),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.PasswordEnvironment),
		"",
		o...,
	)
}
