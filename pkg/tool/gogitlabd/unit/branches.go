package unit

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func branches(name ...string) []*branch.Branch {
	var result []*branch.Branch

	for _, n := range name {
		result = append(
			result,
			branch.New(&gitlab.Branch{Name: n, Commit: &gitlab.Commit{}}),
		)
	}

	return result
}
