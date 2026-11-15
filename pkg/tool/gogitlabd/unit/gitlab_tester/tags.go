package gitlab_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func Tags(name ...string) []*tag.Tag {
	var result []*tag.Tag

	for _, n := range name {
		result = append(result, tag.New(&gitlab.Tag{Name: n}))
	}

	return result
}
