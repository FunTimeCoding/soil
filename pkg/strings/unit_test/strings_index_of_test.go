package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestIndexOf(t *testing.T) {
	assert.Integer(
		t,
		1,
		strings.IndexOf(
			upper.Bravo,
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
		),
	)
}
