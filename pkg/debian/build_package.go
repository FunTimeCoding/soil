package debian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/debian/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func BuildPackage(packageName string) {
	console.Line(
		system.Run(
			constant.DpkgDeb,
			constant.BuildArgument,
			constant.RootOwnerGroup,
			packageName,
		),
	)
}
