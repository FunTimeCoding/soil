package text

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/option"
	"testing"
)

func TestOptimizeWhitespace(t *testing.T) {
	// One blank line is allowed
	assert.String(
		t,
		"",
		OptimizeWhitespace("", nil),
	)
	assert.String(
		t,
		"A\nB\n",
		OptimizeWhitespace("A\nB\n", nil),
	)
	assert.String(
		t,
		"A\n\nB\n",
		OptimizeWhitespace("A\n\nB\n", nil),
	)
	assert.String(
		t,
		"A\n\nB\n",
		OptimizeWhitespace("A\n\n\nB\n", nil),
	)
	assert.String(
		t,
		"A\n\nB\n",
		OptimizeWhitespace("A\n \n \nB\n\n", nil),
	)
	// Fix missing newline at the end
	assert.String(
		t,
		"A\nB\n",
		OptimizeWhitespace("A\nB", nil),
	)
	// No blank line is allowed
	zeroBlank := option.New()
	zeroBlank.AllowedBlankLines = 0
	assert.String(
		t,
		"A\nB\n",
		OptimizeWhitespace("A\n\nB\n", zeroBlank),
	)
}
