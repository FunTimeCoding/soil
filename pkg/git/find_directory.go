package git

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func FindDirectory() string {
	return system.FindDirectoryUp(
		system.WorkDirectory(),
		constant.Directory,
	)
}
