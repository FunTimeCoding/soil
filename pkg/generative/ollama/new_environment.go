package ollama

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment(o ...Option) *Client {
	if s := environment.Optional(constant.OllamaHostEnvironment); s != "" {
		o = append(o, WithHost(s))
	}

	if s := environment.Optional(constant.OllamaPortEnvironment); s != "" {
		o = append(o, WithPort(strings.MustToInteger(s)))
	}

	return New(o...)
}
