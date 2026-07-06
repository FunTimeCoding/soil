package option

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestOutput(t *testing.T) {
	assert.Any(t, &Output{Format: "text"}, New())
	assert.Any(
		t,
		&Output{Format: "text", Debug: true},
		New(WithDebug()),
	)
	assert.Any(
		t,
		&Output{Format: "markdown"},
		New(WithFormat(Markdown)),
	)
	assert.Any(
		t,
		&Output{Format: "notation"},
		New(WithFormat(Notation)),
	)
}
