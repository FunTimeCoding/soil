package pushover

import (
	"github.com/funtimecoding/soil/pkg/pushover/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.UserEnvironment),
		environment.Required(constant.TokenEnvironment),
	)
}
