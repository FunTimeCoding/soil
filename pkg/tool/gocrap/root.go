package gocrap

import (
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/system"
)

func root() string {
	if r := git.FindDirectory(); r != "" {
		return r
	}

	return system.WorkDirectory()
}
