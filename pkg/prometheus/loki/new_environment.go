package loki

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(verbose bool) *Client {
	return New(
		environment.Required(constant.LokiHostEnvironment),
		environment.Required(constant.LokiUserEnvironment),
		environment.Required(constant.LokiPasswordEnvironment),
		verbose,
	)
}
