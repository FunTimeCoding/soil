package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/hypertext"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestTitle(t *testing.T) {
	assert.String(
		t,
		"Test Title",
		hypertext.Title(
			hypertext.Document(
				fixture.File(constant.HypertextPath, "test.html"),
			),
		),
	)
}

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
			hypertext.Document(
				fixture.File(constant.HypertextPath, "test.html"),
			),
		),
	)
}

func TestDivisions(t *testing.T) {
	assert.Any(
		t,
		[]string{"Example DT", "Example DD"},
		hypertext.Divisions(
			hypertext.Document(
				fixture.File(constant.HypertextPath, "test.html"),
			),
		),
	)
}

func TestTables(t *testing.T) {
	assert.Any(
		t,
		[]string{"Example TD"},
		hypertext.Tables(
			hypertext.Document(
				fixture.File(constant.HypertextPath, "test.html"),
			),
		),
	)
}
