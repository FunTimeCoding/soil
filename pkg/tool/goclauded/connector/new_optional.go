package connector

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
)

func NewOptional() *Client {
	if !environment.Exists(constant.TokenEnvironment) {
		return nil
	}

	return NewEnvironment()
}
