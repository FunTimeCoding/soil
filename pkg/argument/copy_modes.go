package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) CopyModes() {
	i.Boolean(
		constant.CopyModes,
		environment.FallbackBoolean(
			provisionConstant.RunnerCopyModesEnvironment,
			false,
		),
		provisionConstant.RunnerCopyModesUsage,
	)
}
