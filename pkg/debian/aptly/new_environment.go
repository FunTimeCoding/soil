package aptly

import (
	"github.com/funtimecoding/soil/pkg/debian/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.RequiredInteger(constant.PortEnvironment),
		environment.Exists(constant.InsecureEnvironment),
		environment.Required(constant.UsernameEnvironment),
		environment.Required(constant.PasswordEnvironment),
	)
}
