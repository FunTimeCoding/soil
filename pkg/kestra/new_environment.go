package kestra

import (
	"github.com/funtimecoding/soil/pkg/kestra/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.TokenEnvironment); s != "" {
		o = append(o, WithToken(s))
	}

	if s := environment.Optional(constant.UserEnvironment); s != "" {
		o = append(o, WithUser(s))
	}

	if s := environment.Optional(constant.PasswordEnvironment); s != "" {
		o = append(o, WithPassword(s))
	}

	return New(environment.Required(constant.HostEnvironment), o...)
}
