package filter

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	assert.NotNil(t, New())
}
