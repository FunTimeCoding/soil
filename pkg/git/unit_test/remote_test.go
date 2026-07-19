package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/remote"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestRemote(t *testing.T) {
	assert.Any(
		t,
		&remote.Remote{
			Name:     "Alfa",
			Locator:  "Bravo",
			Provider: "Charlie",
		},
		remote.New(upper.Alfa, upper.Bravo, upper.Charlie),
	)
}
