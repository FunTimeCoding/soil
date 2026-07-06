package ollama

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.HostEnvironment); s != "" {
		o = append(o, WithHost(s))
	}

	if s := environment.Optional(constant.PortEnvironment); s != "" {
		o = append(o, WithPort(strings.MustToInteger(s)))
	}

	return New(o...)
}
