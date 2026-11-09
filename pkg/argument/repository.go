package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Repository() {
	i.String(
		constant.Repository,
		environment.Fallback(provisionConstant.RunnerRepositoryEnvironment, ""),
		provisionConstant.RunnerRepositoryUsage,
	)
}
