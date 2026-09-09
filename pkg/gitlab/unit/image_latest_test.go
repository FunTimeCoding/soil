package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/image"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestImageLatest(t *testing.T) {
	assert.Any(
		t,
		image.New(&gitlab.RegistryRepositoryTag{Path: "stub:v1.0.2"}),
		image.Latest(
			[]*image.Image{
				image.New(&gitlab.RegistryRepositoryTag{Path: "stub:v1.0.0"}),
				image.New(&gitlab.RegistryRepositoryTag{Path: "stub:v1.0.1"}),
				image.New(&gitlab.RegistryRepositoryTag{Path: "stub:v1.0.2"}),
			},
		),
	)
}
