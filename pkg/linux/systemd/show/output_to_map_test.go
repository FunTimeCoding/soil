package show

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestOutputToMap(t *testing.T) {
	assert.Any(
		t,
		map[string]string{"hello": "world"},
		OutputToMap("hello=world"),
	)
}
