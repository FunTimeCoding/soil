package build

import (
	"github.com/funtimecoding/soil/pkg/strings/contains"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func OutputDirectories() []string {
	var result []string

	for _, d := range system.Subdirectories(constant.Temporary) {
		if !contains.Any(
			constant.SystemArchitectures,
			system.Subdirectories(join.Relative(constant.Temporary, d)),
		) {
			continue
		}

		result = append(result, d)
	}

	return result
}
