package build

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/debian/constant"
	"strings"
)

func DebianPackages() []string {
	var result []string

	for _, d := range system.Files(system.WorkDirectory()) {
		if !strings.HasSuffix(d, constant.PackageExtension) {
			continue
		}

		result = append(result, d)
	}

	return result
}
