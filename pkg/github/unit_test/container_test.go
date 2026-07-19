package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/container"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
)

func TestContainer(t *testing.T) {
	c := container.New(
		&github.Package{
			Name:       new(upper.Alfa),
			Repository: &github.Repository{Name: new(upper.Bravo)},
		},
	)
	c.Raw = nil
	assert.Any(
		t,
		&container.Container{Name: upper.Alfa, Repository: upper.Bravo},
		c,
	)
}
