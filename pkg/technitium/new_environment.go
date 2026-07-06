package technitium

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/technitium/constant"
)

func NewEnvironment() *Client {
	result := New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.TokenEnvironment),
	)

	if environment.Exists(constant.SelfSignedEnvironment) {
		result.SelfSigned()
	}

	return result
}
