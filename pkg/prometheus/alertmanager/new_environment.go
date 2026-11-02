package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.AlertmanagerHostEnvironment),
		environment.FallbackInteger(constant.AlertmanagerPortEnvironment, 0),
		!environment.Exists(constant.AlertmanagerInsecureEnvironment),
		environment.Required(constant.AlertmanagerUserEnvironment),
		environment.Required(constant.AlertmanagerPasswordEnvironment),
	)
}
