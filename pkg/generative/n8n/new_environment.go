package n8n

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.N8nHostEnvironment),
		environment.Required(constant.N8nTokenEnvironment),
	)
}
