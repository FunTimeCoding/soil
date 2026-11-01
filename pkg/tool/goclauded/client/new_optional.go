package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
)

func NewOptional() face.Client {
	if !environment.Exists(constant.TokenEnvironment) {
		return nil
	}

	return NewEnvironment()
}
