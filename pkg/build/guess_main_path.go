package build

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func GuessMainPath(name string) string {
	if s := join.Relative(
		systemConstant.CommandPath,
		name,
		soil.MainFile,
	); system.FileExists(
		s,
	) {
		return s
	}

	return ""
}
