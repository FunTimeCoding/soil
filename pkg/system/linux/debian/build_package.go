package debian

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
)

func BuildPackage(packageName string) {
	console.Line(
		system.Run(
			constant.DpkgDeb,
			constant.DebianBuildArgument,
			constant.DebianRootOwnerGroup,
			packageName,
		),
	)
}
