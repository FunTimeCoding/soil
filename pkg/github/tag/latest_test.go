package tag

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/google/go-github/v88/github"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.String(
		t,
		"v1.0.1",
		Latest(
			[]*github.RepositoryTag{
				{Name: new("v1.0.0")},
				{Name: new("v1.0.1")},
			},
		).GetName(),
	)
}
