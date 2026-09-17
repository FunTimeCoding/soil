package lint

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/system"
)

func Root(flag string) (string, string) {
	work := system.WorkDirectory()

	if flag != "" {
		root := system.AbsolutePath(flag)

		return root, root
	}

	root := git.FindDirectory()

	if root == "" {
		system.Exitf(
			1,
			"no repository found above %s, pass --%s\n",
			work,
			constant.Root,
		)
	}

	return root, work
}
