package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/image"
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
		image.Sort(
			[]*gitlab.RegistryRepositoryTag{
				{Path: "stub:v1.0.0"},
				{Path: "stub:v1.0.1"},
				{Path: "stub:v1.0.2"},
			},
		),
	)
}
