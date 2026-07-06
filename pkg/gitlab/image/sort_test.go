package image

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestSort(t *testing.T) {
	assert.Any(
		t,
		[]*gitlab.RegistryRepositoryTag{
			{Path: "stub:v1.0.2"},
			{Path: "stub:v1.0.1"},
			{Path: "stub:v1.0.0"},
		},
		Sort(
			[]*gitlab.RegistryRepositoryTag{
				{Path: "stub:v1.0.0"},
				{Path: "stub:v1.0.1"},
				{Path: "stub:v1.0.2"},
			},
		),
	)
}
