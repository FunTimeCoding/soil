package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func MergeRequest() {
	g := gitlab.NewEnvironment()
	defer g.Close()
	f := constant.Format

	for _, r := range g.MustProjectsMergeRequests(context.Background()) {
		console.Line(r.Format(f))
	}
}
