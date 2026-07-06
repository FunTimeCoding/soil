package semver

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"strings"
)

func Trim(version string) string {
	return strings.TrimPrefix(version, constant.VersionPrefix)
}
