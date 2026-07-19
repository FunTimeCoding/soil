package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/description"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestDescription(t *testing.T) {
	assert.Any(
		t,
		&description.Description{Title: "Alfa", Short: "Bravo"},
		description.New(upper.Alfa, upper.Bravo),
	)
	assert.Any(
		t,
		&description.Description{Type: "Alfa", Title: "Bravo", Short: "Charlie"},
		description.NewType(upper.Alfa, upper.Bravo, upper.Charlie),
	)
}
