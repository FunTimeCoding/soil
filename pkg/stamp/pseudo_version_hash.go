package stamp

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/stamp/constant"
	"golang.org/x/mod/module"
	"golang.org/x/mod/semver"
	"strings"
)

func pseudoVersionHash(version string) string {
	v := strings.TrimSuffix(version, semver.Build(version))

	if !module.IsPseudoVersion(v) {
		return constant.DefaultGitHash
	}

	r, e := module.PseudoVersionRev(v)

	if e != nil || len(r) < gitConstant.HashLength {
		return constant.DefaultGitHash
	}

	return r[:gitConstant.HashLength]
}
