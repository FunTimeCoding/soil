package pricing

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFormatTokens(t *testing.T) {
	assert.String(t, "156", FormatTokens(156))
	assert.String(t, "23.7K", FormatTokens(23727))
	assert.String(t, "1.5M", FormatTokens(1514005))
	assert.String(t, "53.5M", FormatTokens(53534614))
	assert.String(t, "3.1B", FormatTokens(3113800000))
}
