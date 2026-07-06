package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func File() {
	a := argument.NewSimple("gitlab-file")
	a.String(argument.Owner, "", "group or user")
	a.String(argument.Repository, "", "repository name")
	a.String(argument.Branch, "", "branch name")
	a.String(argument.File, "", "file path")
	a.ParseSimple()
	owner := a.GetString(argument.Owner)
	repository := a.GetString(argument.Repository)
	branch := a.GetString(argument.Branch)
	file := a.GetString(argument.File)
	g := gitlab.NewEnvironment()
	p := g.MustProjectByName(owner, repository)
	f := g.MustFile(p.Identifier, branch, file)
	fmt.Printf("File: %+v\n", f)
}
