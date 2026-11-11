package backrest

import (
	"github.com/funtimecoding/soil/pkg/backrest/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	if c := environment.Slice(constant.CommandEnvironment); len(c) > 0 {
		return New(c...)
	}

	return New(constant.Executable)
}
