package n8n

import (
	"github.com/funtimecoding/soil/pkg/generative/n8n/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.TokenEnvironment),
	)
}
