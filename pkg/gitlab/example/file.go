package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func File() {
	a := argument.NewSimple("gitlab-file")
	a.String(argumentConstant.Owner, "", "group or user")
	a.String(argumentConstant.Repository, "", "repository name")
	a.String(argumentConstant.Branch, "", "branch name")
	a.String(argumentConstant.File, "", "file path")
	a.ParseSimple()
	owner := a.GetString(argumentConstant.Owner)
	repository := a.GetString(argumentConstant.Repository)
	branch := a.GetString(argumentConstant.Branch)
	file := a.GetString(argumentConstant.File)
	g := gitlab.NewEnvironment()
	defer g.Close()
	x := context.Background()
	p := g.MustProjectByName(x, owner, repository)
	f := g.MustFile(x, p.Identifier, branch, file)
	console.Format("File: %+v\n", f)
}
