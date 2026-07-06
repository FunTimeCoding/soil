package tag

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.Any(
		t,
		&gitlab.Tag{Name: "v1.0.2"},
		Latest(
			[]*gitlab.Tag{
				{Name: "v1.0.0"},
				{Name: "v1.0.2"},
				{Name: "v1.0.1"},
			},
		),
	)
}
