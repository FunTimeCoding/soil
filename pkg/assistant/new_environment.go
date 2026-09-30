package assistant

import (
	"github.com/funtimecoding/soil/pkg/assistant/constant"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(
	l *logger.Logger,
	r face.Reporter,
	o ...Option,
) *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.TokenEnvironment),
		l,
		r,
		o...,
	)
}
