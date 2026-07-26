package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func BranchRequest() {
	a := argument.NewSimple("gitlab-branch-request")
	a.Integer64(argumentConstant.Project, 0, "Project ID")
	a.String(argumentConstant.Branch, "", "Branch name")
	a.ParseSimple()
	g := gitlab.NewEnvironment()
	f := constant.Format
	fmt.Println(
		g.MustBranchRequest(
			a.GetInteger64(argumentConstant.Project),
			a.GetString(argumentConstant.Branch),
		).Format(f),
	)
}
