package action

import (
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func IsActionRun() bool {
	return environment.Exists(constant.RunEnvironment)
}
