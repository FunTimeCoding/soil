package remote

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestRemote(t *testing.T) {
	assert.Any(
		t,
		&Remote{
			Name:     "Alfa",
			Locator:  "Bravo",
			Provider: "Charlie",
		},
		New(upper.Alfa, upper.Bravo, upper.Charlie),
	)
}
