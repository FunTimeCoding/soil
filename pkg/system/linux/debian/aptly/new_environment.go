package aptly

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.AptlyHostEnvironment),
		environment.RequiredInteger(constant.AptlyPortEnvironment),
		environment.Exists(constant.AptlyInsecureEnvironment),
		environment.Required(constant.AptlyUsernameEnvironment),
		environment.Required(constant.AptlyPasswordEnvironment),
	)
}
