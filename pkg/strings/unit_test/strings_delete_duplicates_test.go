package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	assert.Any(
		t,
		[]string{upper.Alfa, upper.Bravo},
		strings.DeleteDuplicates([]string{upper.Alfa, upper.Alfa, upper.Bravo}),
	)
}
