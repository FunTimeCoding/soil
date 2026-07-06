package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/project"
)

func Search() {
	g := gitlab.NewEnvironment()
	// Free version search is limited
	for _, p := range g.MustSearchProject("") {
		fmt.Printf("Project: %s\n", p.Raw.NameWithNamespace)
	}

	for _, b := range g.MustSearchBlob(
		fmt.Sprintf("filename:%s", project.GitLabFile),
	) {
		fmt.Printf("Blob: %+v\n", b)
	}
}
