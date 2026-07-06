package ranges

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestRange(t *testing.T) {
	assert.Any(t, &Range{L: 1, R: 2}, New(1, 2))
	assert.Any(t, &Range{L: -1, R: 3}, FromFactor(1, 2))
}
