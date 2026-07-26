package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/constant"
)

func BranchRequest() {
	a := argument.NewSimple("github-branch-request")
	a.String(argumentConstant.Branch, "", "Branch name")
	a.ParseSimple()
	branch := a.GetString(argumentConstant.Branch)
	c := github.NewEnvironment()
	f := option.ExtendedColor.Copy()
	fmt.Println(
		c.MustBranchRequest(
			constant.Namespace,
			constant.Repository,
			branch,
		).Format(f),
	)
}
