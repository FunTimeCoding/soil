package salt

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/constant"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	var o []Option

	if v := environment.Optional(
		constant.AuthenticationEnvironment,
	); v != "" {
		o = append(o, WithAuthentication(v))
	}

	return New(
		environment.Required(constant.HostEnvironment),
		strings.MustToInteger(
			environment.Fallback(constant.PortEnvironment, "8000"),
		),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.PasswordEnvironment),
		o...,
	)
}
