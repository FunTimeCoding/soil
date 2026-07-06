package grafana

import (
	"github.com/funtimecoding/soil/pkg/grafana/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.PortEnvironment),
		environment.Required(constant.TokenEnvironment),
	)
}
