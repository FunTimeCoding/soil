package build

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func GuessMainPath(name string) string {
	if s := join.Relative(
		constant.CommandPath,
		name,
		library.MainFile,
	); system.FileExists(s) {
		return s
	}

	return ""
}
