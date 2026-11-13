package repository

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func InGitDirectory(path string) bool {
	return strings.Contains(
		join.Empty(stringsConstant.Slash, path, stringsConstant.Slash),
		join.Empty(
			stringsConstant.Slash,
			gitConstant.Directory,
			stringsConstant.Slash,
		),
	)
}
