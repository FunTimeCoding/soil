package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.KeyEnvironment),
		environment.Required(constant.SecretEnvironment),
		environment.Exists(constant.InsecureEnvironment),
	)
}
