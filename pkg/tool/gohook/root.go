package gohook

import (
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/terminal"
)

func root(t *terminal.Terminal) string {
	r := git.FindDirectory()

	if r == "" {
		t.Exitf("no repository found above %s\n", system.WorkDirectory())
	}

	return r
}
