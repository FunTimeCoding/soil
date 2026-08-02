package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/container"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/google/go-github/v89/github"
	"testing"
)

func TestContainer(t *testing.T) {
	c := container.New(
		&github.Package{
			Name:       new(constant.UpperAlfa),
			Repository: &github.Repository{Name: new(constant.UpperBravo)},
		},
	)
	c.Raw = nil
	assert.Any(
		t,
		&container.Container{
			Name:       "Alfa",
			Repository: "Bravo",
		},
		c,
	)
}
