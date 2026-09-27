package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/build/origin"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func ModuleOrigin(
	module string,
	version string,
) *origin.Origin {
	r := run.New()
	r.Panic = false

	return ParseOrigin(
		r.Start(
			constant.Go,
			constant.List,
			constant.ModuleArgument,
			constant.NotationArgument,
			fmt.Sprintf("%s@%s", module, version),
		),
	)
}
