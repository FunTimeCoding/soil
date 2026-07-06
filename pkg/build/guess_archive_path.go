package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func GuessArchivePath(
	name string,
	systemArchitecture string,
) string {
	if s := join.Relative(
		constant.Temporary,
		fmt.Sprintf("%s-%s.zip", name, systemArchitecture),
	); system.FileExists(s) {
		return s
	}

	return ""
}
