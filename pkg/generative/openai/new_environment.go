package openai

import (
	"github.com/funtimecoding/soil/pkg/generative/openai/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.TokenEnvironment))
}
