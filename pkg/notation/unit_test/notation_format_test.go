package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/notation/fixture"
	"testing"
)

func TestFormatTypes(t *testing.T) {
	assert.Any(
		t,
		"{\n    \"String\": \"a\",\n    \"Integer\": 1,\n    \"Float\": 1.5,\n    \"Boolean\": true\n}",
		notation.Format(
			fixture.Primitives{
				String:  "a",
				Integer: 1,
				Float:   1.5,
				Boolean: true,
			},
		),
	)
}

func TestFormatStringWithVector(t *testing.T) {
	assert.Any(
		t,
		"{\n    \"String\": \"1,<1.0, 1.0, 1.0>,2\"\n}",
		notation.Format(
			fixture.WithString{
				String: "1,<1.0, 1.0, 1.0>,2",
			},
		),
	)
}
