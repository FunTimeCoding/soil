package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestProject(t *testing.T) {
	p := project.New(&gitlab.Project{Namespace: &gitlab.ProjectNamespace{}})
	p.Raw = nil
	assert.Exported(t, &project.Project{}, p)
}
