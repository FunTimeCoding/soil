package image

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.Any(
		t,
		&gitlab.RegistryRepositoryTag{Path: "stub:v1.0.2"},
		Latest(
			[]*gitlab.RegistryRepositoryTag{
				{Path: "stub:v1.0.0"},
				{Path: "stub:v1.0.1"},
				{Path: "stub:v1.0.2"},
			},
		),
	)
}
