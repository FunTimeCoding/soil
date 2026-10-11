package packager

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func binaryRoot(packageDirectory string) string {
	return join.Absolute(
		packageDirectory,
		constant.Resources,
		constant.Local,
		constant.Binary,
	)
}
