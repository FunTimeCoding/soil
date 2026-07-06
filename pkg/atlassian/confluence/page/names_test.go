package page

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestNames(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		Names([]*Page{{Name: upper.Alfa}, {Name: upper.Bravo}}),
	)
}
