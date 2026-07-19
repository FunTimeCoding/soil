package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/hypertext"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestHeaders(t *testing.T) {
	assert.Any(
		t,
		[]string{
			"Example h1",
			"Example h2",
			"Example h3",
			"Example h4",
			"Example h5",
			"Example h6",
		},
		hypertext.Headers(
			hypertext.Document(fixture.File(constant.HypertextPath, "test.html")),
		),
	)
}
