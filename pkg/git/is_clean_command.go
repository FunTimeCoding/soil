package git

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func IsCleanCommand() bool {
	return system.Run(
		constant.Command,
		constant.Status,
		constant.Porcelain,
	) == ""
}
