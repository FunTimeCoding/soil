package example

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/constant"
)

func BranchRequest() {
	a := argument.NewSimple("github-branch-request")
	a.String(argumentConstant.Branch, "", "Branch name")
	a.ParseSimple()
	branch := a.GetString(argumentConstant.Branch)
	c := github.NewEnvironment()
	f := consoleConstant.ExtendedColorFormat.Copy()
	console.Line(
		c.MustBranchRequest(
			constant.Namespace,
			constant.Repository,
			branch,
		).Format(
			f,
		),
	)
}
