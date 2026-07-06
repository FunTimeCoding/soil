package hypertext

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestDivisions(t *testing.T) {
	assert.Any(
		t,
		[]string{"Example DT", "Example DD"},
		Divisions(
			Document(fixture.File(constant.HypertextPath, "test.html")),
		),
	)
}
