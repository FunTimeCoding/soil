package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func MergeRequest() {
	g := gitlab.NewEnvironment()
	f := constant.Format

	for _, r := range g.ProjectsMergeRequests() {
		fmt.Println(r.Format(f))
	}
}
