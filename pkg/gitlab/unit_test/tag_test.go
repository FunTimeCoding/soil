package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestTagLatest(t *testing.T) {
	assert.Any(
		t,
		&gitlab.Tag{Name: "v1.0.2"},
		tag.Latest(
			[]*gitlab.Tag{
				{Name: "v1.0.0"},
				{Name: "v1.0.2"},
				{Name: "v1.0.1"},
			},
		),
	)
}
