package unit

import (
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func seedTarget() *project.Project {
	return project.New(&gitlab.Project{Namespace: &gitlab.ProjectNamespace{}})
}
