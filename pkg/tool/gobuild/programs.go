package gobuild

import (
	"github.com/funtimecoding/soil/pkg/build/constant"
	"github.com/funtimecoding/soil/pkg/system"
	constant1 "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func programs() []string {
	var result []string

	for _, n := range system.Directories(
		join.Absolute(system.WorkDirectory(), constant1.CommandPath),
	) {
		if n != constant.ExamplePath {
			result = append(result, n)
		}
	}

	return result
}
