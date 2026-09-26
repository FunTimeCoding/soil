package semver

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func Prefix(version string) string {
	if strings.HasPrefix(version, constant.VersionPrefix) {
		return version
	}

	return join.Empty(constant.VersionPrefix, version)
}
