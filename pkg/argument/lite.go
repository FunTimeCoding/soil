package argument

import (
	"github.com/funtimecoding/soil/pkg/relational/lite/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Lite() {
	i.String(
		Lite,
		environment.Fallback(
			constant.PathEnvironment,
			i.identity.LitePath(),
		),
		constant.PathUsage,
	)
}
