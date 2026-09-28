package fritz

import (
	"github.com/funtimecoding/soil/pkg/fritz/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Fallback(constant.HostEnvironment, constant.DefaultHost),
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.PasswordEnvironment),
	)
}
