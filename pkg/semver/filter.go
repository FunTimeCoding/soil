package semver

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"strings"
)

func Filter(tags []string) []string {
	var result []string

	for _, t := range tags {
		if strings.HasPrefix(t, gitConstant.VersionPrefix) {
			result = append(result, t)
		}
	}

	return result
}
