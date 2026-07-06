package git

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"strings"
)

func Versions(path string) semver.Versions {
	var result semver.Versions

	for _, t := range Tags(path) {
		if strings.HasPrefix(t, constant.VersionPrefix) {
			result = append(
				result,
				semver.New(strings.TrimPrefix(t, constant.VersionPrefix)),
			)
		}
	}

	semver.Sort(result)

	return result
}
