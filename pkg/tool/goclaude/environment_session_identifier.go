package goclaude

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
)

func environmentSessionIdentifier() string {
	if i := environment.Optional(constant.SessionIdentifierEnvironment); i != "" {
		return i
	}

	return readHookInput().SessionIdentifier
}
