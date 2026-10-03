package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text"
	"github.com/funtimecoding/soil/pkg/text/option"
	"testing"
)

func TestOptimizeWhitespaceKeepsOneBlankLine(t *testing.T) {
	assert.String(t, "", text.OptimizeWhitespace("", nil))
	assert.String(t, "A\nB\n", text.OptimizeWhitespace("A\nB\n", nil))
	assert.String(t, "A\n\nB\n", text.OptimizeWhitespace("A\n\nB\n", nil))
	assert.String(t, "A\n\nB\n", text.OptimizeWhitespace("A\n\n\nB\n", nil))
	assert.String(t, "A\n\nB\n", text.OptimizeWhitespace("A\n \n \nB\n\n", nil))
}

func TestOptimizeWhitespaceAddsMissingFinalNewline(t *testing.T) {
	assert.String(t, "A\nB\n", text.OptimizeWhitespace("A\nB", nil))
}

func TestOptimizeWhitespaceRemovesBlankLinesWhenNoneAllowed(t *testing.T) {
	zeroBlank := option.New()
	zeroBlank.AllowedBlankLines = 0
	assert.String(t, "A\nB\n", text.OptimizeWhitespace("A\n\nB\n", zeroBlank))
}
