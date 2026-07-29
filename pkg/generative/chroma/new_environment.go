package chroma

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.ChromaHostEnvironment),
		strings.MustToInteger(
			environment.Required(constant.ChromaPortEnvironment),
		),
	)
}
