package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func Search() {
	g := gitlab.NewEnvironment()
	// Free version search is limited
	for _, p := range g.MustSearchProject("") {
		fmt.Printf("Project: %s\n", p.Raw.NameWithNamespace)
	}

	for _, b := range g.MustSearchBlob(
		fmt.Sprintf("filename:%s", constant.GitLabFile),
	) {
		fmt.Printf("Blob: %+v\n", b)
	}
}
