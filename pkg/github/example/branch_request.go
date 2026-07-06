package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/constant"
)

func BranchRequest() {
	a := argument.NewSimple("github-branch-request")
	a.String(argument.Branch, "", "Branch name")
	a.ParseSimple()
	branch := a.GetString(argument.Branch)
	c := github.NewEnvironment()
	f := option.ExtendedColor.Copy()
	fmt.Println(
		c.MustBranchRequest(
			constant.LibraryNamespace,
			constant.LibraryRepository,
			branch,
		).Format(f),
	)
}
