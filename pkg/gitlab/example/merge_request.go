package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func MergeRequest() {
	g := gitlab.NewEnvironment()
	f := constant.Format

	for _, r := range g.MustProjectsMergeRequests() {
		console.Line(r.Format(f))
	}
}
