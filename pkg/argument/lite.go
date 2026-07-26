package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	relational "github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Lite() {
	i.String(
		constant.Lite,
		environment.Fallback(
			relational.LitePathEnvironment,
			i.identity.LitePath(),
		),
		relational.LitePathUsage,
	)
}
