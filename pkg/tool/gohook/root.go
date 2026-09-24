package gohook

import (
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/system"
)

func root() string {
	r := git.FindDirectory()

	if r == "" {
		system.Exitf(
			1,
			"no repository found above %s\n",
			system.WorkDirectory(),
		)
	}

	return r
}
