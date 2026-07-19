package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMapHasKey(t *testing.T) {
	m := map[string]any{"height": 180}
	assert.MapHasKey(t, m, "height")
}
