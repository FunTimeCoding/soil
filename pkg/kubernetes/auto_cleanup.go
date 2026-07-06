package kubernetes

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func AutoCleanup() bool {
	return environment.Exists(constant.AutoCleanupEnvironment)
}
