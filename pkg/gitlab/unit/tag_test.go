package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"testing"
)

func TestTagLatest(t *testing.T) {
	assert.Any(
		t,
		&tag.Tag{Name: "v1.0.2"},
		tag.Latest(
			[]*tag.Tag{{Name: "v1.0.0"}, {Name: "v1.0.2"}, {Name: "v1.0.1"}},
		),
	)
}
