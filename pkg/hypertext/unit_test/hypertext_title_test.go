package unit_test

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
				fixture.File(
					constant.HypertextPath,
					"test.html",
				),
			),
		),
	)
}
