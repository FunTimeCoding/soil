package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/maps"
	"testing"
)

func TestToMarkup(t *testing.T) {
	assert.String(
		t,
		"  a: dGVzdA==\n",
		maps.ToMarkup(map[string]string{"a": "test"}),
	)
}
