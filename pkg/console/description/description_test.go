package description

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestDescription(t *testing.T) {
	assert.Any(
		t,
		&Description{Title: "Alfa", Short: "Bravo"},
		New(upper.Alfa, upper.Bravo),
	)
	assert.Any(
		t,
		&Description{Type: "Alfa", Title: "Bravo", Short: "Charlie"},
		NewType(upper.Alfa, upper.Bravo, upper.Charlie),
	)
}
