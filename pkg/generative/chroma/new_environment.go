package chroma

import (
	"github.com/funtimecoding/soil/pkg/generative/chroma/constant"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		strings.MustToInteger(
			environment.Required(constant.PortEnvironment),
		),
	)
}
