package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	goclaudeConstant "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
)

func NewOptional() face.Client {
	if !environment.Exists(goclaudeConstant.TokenEnvironment) {
		return nil
	}

	return NewEnvironment()
}
