package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestToMap(t *testing.T) {
	assert.Any(
		t,
		map[string]string{"A": "a", "B": "b"},
		strings.ToMap([]string{"A=a", "B=b"}, "="),
	)
}
