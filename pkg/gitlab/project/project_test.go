package project

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestProject(t *testing.T) {
	p := New(&gitlab.Project{Namespace: &gitlab.ProjectNamespace{}})
	p.Raw = nil
	assert.Any(t, &Project{}, p)
}
