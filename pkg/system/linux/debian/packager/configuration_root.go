package packager

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func configurationRoot(packageDirectory string) string {
	return join.Absolute(
		packageDirectory,
		constant.DebianPackageConfigurationDirectory,
	)
}
