package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"testing"
)

func TestTrimSuffix(t *testing.T) {
	assert.Strings(
		t,
		[]string{"a", "b"},
		slice.TrimSuffix([]string{"a.", "b."}, separator.Dot),
	)
}
