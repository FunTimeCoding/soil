package band

import (
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.AddressEnvironment),
		environment.Fallback(constant.UserEnvironment, constant.DefaultUser),
		environment.Required(constant.PasswordEnvironment),
	)
}
