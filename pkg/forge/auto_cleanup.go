package forge

import (
	"github.com/funtimecoding/soil/pkg/forge/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func AutoCleanup() bool {
	return environment.Exists(constant.AutoCleanupEnvironment)
}
