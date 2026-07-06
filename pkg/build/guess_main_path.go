package build

import (
	"github.com/funtimecoding/soil/pkg/project"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func GuessMainPath(name string) string {
	if s := join.Relative(
		constant.CommandPath,
		name,
		project.MainFile,
	); system.FileExists(s) {
		return s
	}

	return ""
}
