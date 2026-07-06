package container

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v88/github"
	"testing"
)

func TestContainer(t *testing.T) {
	c := New(
		&github.Package{
			Name:       new(upper.Alfa),
			Repository: &github.Repository{Name: new(upper.Bravo)},
		},
	)
	c.Raw = nil
	assert.Any(
		t,
		&Container{Name: upper.Alfa, Repository: upper.Bravo},
		c,
	)
}
