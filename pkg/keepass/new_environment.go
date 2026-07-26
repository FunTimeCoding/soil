package keepass

import (
	"github.com/funtimecoding/soil/pkg/keepass/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.DatabaseEnvironment),
		environment.Required(constant.PasswordEnvironment),
	)
}
