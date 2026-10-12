package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	relationalConstant "github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Lite() {
	i.String(
		constant.Lite,
		environment.Fallback(
			relationalConstant.LitePathEnvironment,
			i.identity.LitePath(),
		),
		relationalConstant.LitePathUsage,
	)
}
