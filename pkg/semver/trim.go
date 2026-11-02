package semver

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"strings"
)

func Trim(version string) string {
	return strings.TrimPrefix(version, gitConstant.VersionPrefix)
}
