package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	lite "github.com/funtimecoding/soil/pkg/relational/lite/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Lite() {
	i.String(
		constant.Lite,
		environment.Fallback(
			lite.PathEnvironment,
			i.identity.LitePath(),
		),
		lite.PathUsage,
	)
}
