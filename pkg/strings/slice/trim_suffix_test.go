package slice

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"testing"
)

func TestTrimSuffix(t *testing.T) {
	assert.Strings(
		t,
		[]string{"a", "b"},
		TrimSuffix([]string{"a.", "b."}, separator.Dot),
	)
}
