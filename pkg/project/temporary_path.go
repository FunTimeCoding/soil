package project

import (
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func TemporaryPath(path ...string) string {
	return join.Absolute(
		git.FindDirectory(),
		constant.Temporary,
		join.Relative(path...),
	)
}
