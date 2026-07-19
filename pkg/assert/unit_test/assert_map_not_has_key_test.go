package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMapNotHasKey(t *testing.T) {
	m := map[string]any{"height": 180}
	assert.MapNotHasKey(t, m, "weight")
}
