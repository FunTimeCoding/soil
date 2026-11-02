package example

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func Search() {
	g := gitlab.NewEnvironment()
	defer g.Close()
	x := context.Background()

	for _, p := range g.MustSearchProject(x, "") {
		console.Format("Project: %s\n", p.Raw.NameWithNamespace)
	}

	for _, b := range g.MustSearchBlob(
		x,
		fmt.Sprintf("filename:%s", soil.GitLabFile),
	) {
		console.Format("Blob: %+v\n", b)
	}
}
