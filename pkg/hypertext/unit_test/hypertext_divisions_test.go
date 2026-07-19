package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/hypertext"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestDivisions(t *testing.T) {
	assert.Any(
		t,
		[]string{"Example DT", "Example DD"},
		hypertext.Divisions(
			hypertext.Document(fixture.File(constant.HypertextPath, "test.html")),
		),
	)
}
