package assistant

import (
	"github.com/funtimecoding/soil/pkg/assistant/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.TokenEnvironment),
		o...,
	)
}
