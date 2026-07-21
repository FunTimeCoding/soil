package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"testing"
)

func TestFormatTokens(t *testing.T) {
	assert.String(t, "156", pricing.FormatTokens(156))
	assert.String(t, "23.7K", pricing.FormatTokens(23727))
	assert.String(t, "1.5M", pricing.FormatTokens(1514005))
	assert.String(t, "53.5M", pricing.FormatTokens(53534614))
	assert.String(t, "3.1B", pricing.FormatTokens(3113800000))
}
