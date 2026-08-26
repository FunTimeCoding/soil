package debian

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/debian/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func BuildPackage(packageName string) {
	fmt.Println(
		system.Run(
			constant.DpkgDeb,
			constant.BuildArgument,
			constant.RootOwnerGroup,
			packageName,
		),
	)
}
