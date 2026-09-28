package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.RequiredInteger(constant.PortEnvironment),
		environment.Required(constant.TokenEnvironment),
	)
}
