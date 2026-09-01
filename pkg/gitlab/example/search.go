package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func Search() {
	g := gitlab.NewEnvironment()
	// Free version search is limited
	for _, p := range g.MustSearchProject("") {
		console.Format("Project: %s\n", p.Raw.NameWithNamespace)
	}

	for _, b := range g.MustSearchBlob(
		fmt.Sprintf("filename:%s", constant.GitLabFile),
	) {
		console.Format("Blob: %+v\n", b)
	}
}
