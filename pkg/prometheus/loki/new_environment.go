package loki

import (
	"github.com/funtimecoding/soil/pkg/prometheus/loki/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(verbose bool) *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.PasswordEnvironment),
		verbose,
	)
}
