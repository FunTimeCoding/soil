package build

import (
	"github.com/funtimecoding/soil/pkg/strings/contains"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func OutputDirectories() []string {
	var result []string

	for _, d := range system.Subdirectories(systemConstant.Temporary) {
		if !contains.Any(
			systemConstant.SystemArchitectures,
			system.Subdirectories(join.Relative(systemConstant.Temporary, d)),
		) {
			continue
		}

		result = append(result, d)
	}

	return result
}
