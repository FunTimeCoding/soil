package gobuild

import (
	buildConstant "github.com/funtimecoding/soil/pkg/build/constant"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func programs() []string {
	var result []string

	for _, n := range system.Directories(
		join.Absolute(system.WorkDirectory(), systemConstant.CommandPath),
	) {
		if n != buildConstant.ExamplePath {
			result = append(result, n)
		}
	}

	return result
}
