package semver

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func Prefix(version string) string {
	if strings.HasPrefix(version, gitConstant.VersionPrefix) {
		return version
	}

	return join.Empty(gitConstant.VersionPrefix, version)
}
