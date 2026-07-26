package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func File() {
	a := argument.NewSimple("gitlab-file")
	a.String(constant.Owner, "", "group or user")
	a.String(constant.Repository, "", "repository name")
	a.String(constant.Branch, "", "branch name")
	a.String(constant.File, "", "file path")
	a.ParseSimple()
	owner := a.GetString(constant.Owner)
	repository := a.GetString(constant.Repository)
	branch := a.GetString(constant.Branch)
	file := a.GetString(constant.File)
	g := gitlab.NewEnvironment()
	p := g.MustProjectByName(owner, repository)
	f := g.MustFile(p.Identifier, branch, file)
	fmt.Printf("File: %+v\n", f)
}
