package openbao

import (
	"github.com/funtimecoding/soil/pkg/openbao/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewEnvironment() *Client {
	l := locator.New(environment.Required(constant.HostEnvironment)).Port(
		environment.FallbackInteger(constant.PortEnvironment, 0),
	)

	if environment.Exists(constant.InsecureEnvironment) {
		l.Insecure()
	}

	result := New(l)

	if r := environment.Optional(constant.RoleEnvironment); r != "" {
		result.WithKubernetes(r, constant.TokenPath)
	}

	return result
}
