package build

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func GuessBinaryPath(
	name string,
	systemArchitecture string,
) string {
	if s := join.Relative(
		constant.Temporary,
		name,
		systemArchitecture,
		name,
	); system.FileExists(s) {
		return s
	}

	return ""
}
