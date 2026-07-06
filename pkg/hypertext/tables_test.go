package hypertext

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestTables(t *testing.T) {
	assert.Any(
		t,
		[]string{"Example TD"},
		Tables(
			Document(fixture.File(constant.HypertextPath, "test.html")),
		),
	)
}
