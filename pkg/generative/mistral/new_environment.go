package mistral

import (
	"github.com/funtimecoding/soil/pkg/generative/mistral/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.TokenEnvironment))
}
