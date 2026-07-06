package file

import (
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func collect(paths []string) []string {
	result := paths

	if s := environment.Optional(constant.FileEnvironment); s != "" {
		result = append(result, split.Comma(s)...)
	}

	return result
}
