package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/tag"
	"github.com/google/go-github/v89/github"
	"testing"
)

func TestTagLatest(t *testing.T) {
	assert.String(
		t,
		"v1.0.1",
		tag.Latest(
			[]*github.RepositoryTag{
				{Name: new("v1.0.0")},
				{Name: new("v1.0.1")},
			},
		).GetName(),
	)
}
