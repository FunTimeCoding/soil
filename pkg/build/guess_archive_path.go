package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func GuessArchivePath(
	name string,
	systemArchitecture string,
) string {
	if s := join.Relative(
		systemConstant.Temporary,
		fmt.Sprintf("%s-%s.zip", name, systemArchitecture),
	); system.FileExists(s) {
		return s
	}

	return ""
}
