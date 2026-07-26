package salt

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	var o []Option

	if v := environment.Optional(
		constant.SaltAuthenticationEnvironment,
	); v != "" {
		o = append(o, WithAuthentication(v))
	}

	return New(
		environment.Required(constant.SaltHostEnvironment),
		strings.MustToInteger(
			environment.Fallback(constant.SaltPortEnvironment, "8000"),
		),
		environment.Required(constant.SaltUserEnvironment),
		environment.Required(constant.SaltPasswordEnvironment),
		o...,
	)
}
