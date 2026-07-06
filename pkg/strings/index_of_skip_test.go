package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestIndexOfSkip(t *testing.T) {
	assert.Integer(
		t,
		1,
		IndexOfSkip(
			upper.Bravo,
			[]string{upper.Alfa, upper.Bravo, upper.Charlie, upper.Bravo},
			0,
		),
	)
	assert.Integer(
		t,
		3,
		IndexOfSkip(
			upper.Bravo,
			[]string{upper.Alfa, upper.Bravo, upper.Charlie, upper.Bravo},
			1,
		),
	)
}
