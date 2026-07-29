package build

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"strings"
)

func DebianPackages() []string {
	var result []string

	for _, d := range system.Files(system.WorkDirectory()) {
		if !strings.HasSuffix(d, constant.DebianPackageExtension) {
			continue
		}

		result = append(result, d)
	}

	return result
}
