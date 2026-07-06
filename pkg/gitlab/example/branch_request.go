package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func BranchRequest() {
	a := argument.NewSimple("gitlab-branch-request")
	a.Integer64(argument.Project, 0, "Project ID")
	a.String(argument.Branch, "", "Branch name")
	a.ParseSimple()
	g := gitlab.NewEnvironment()
	f := constant.Format
	fmt.Println(
		g.MustBranchRequest(
			a.GetInteger64(argument.Project),
			a.GetString(argument.Branch),
		).Format(f),
	)
}
