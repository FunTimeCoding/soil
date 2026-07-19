package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMapValue(t *testing.T) {
	m := map[string]any{"height": 180, "name": "Fern"}
	assert.MapValue(t, 180, m, "height")
	assert.MapValue(t, "Fern", m, "name")
}
