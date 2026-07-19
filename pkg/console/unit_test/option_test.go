package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/option"
	"testing"
)

func TestOutput(t *testing.T) {
	assert.Any(t, &option.Output{Format: "text"}, option.New())
	assert.Any(
		t,
		&option.Output{Format: "text", Debug: true},
		option.New(option.WithDebug()),
	)
	assert.Any(
		t,
		&option.Output{Format: "markdown"},
		option.New(option.WithFormat(option.Markdown)),
	)
	assert.Any(
		t,
		&option.Output{Format: "notation"},
		option.New(option.WithFormat(option.Notation)),
	)
}
