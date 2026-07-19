package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestDeleteEmpty(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		strings.DeleteEmpty([]string{"", upper.Alfa, "", upper.Bravo, ""}),
	)
}
