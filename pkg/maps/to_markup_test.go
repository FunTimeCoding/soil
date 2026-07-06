package maps

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestToMarkup(t *testing.T) {
	assert.String(
		t,
		"  a: dGVzdA==\n",
		ToMarkup(map[string]string{"a": "test"}),
	)
}
