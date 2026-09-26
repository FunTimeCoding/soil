package aptly

import (
	"github.com/funtimecoding/soil/pkg/debian/constant"
	"strings"
)

func PackageVersions(
	packages []string,
	name string,
) []string {
	var result []string

	for _, p := range packages {
		fields := strings.Fields(p)

		if len(fields) < constant.PackageKeyFields || fields[1] != name {
			continue
		}

		result = append(result, fields[2])
	}

	return result
}
