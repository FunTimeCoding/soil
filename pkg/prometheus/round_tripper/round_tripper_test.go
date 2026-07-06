package round_tripper

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestRoundTripper(t *testing.T) {
	assert.NotNil(t, New("", ""))
}
