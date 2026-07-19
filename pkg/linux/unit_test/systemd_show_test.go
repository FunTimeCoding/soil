package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/linux/systemd/show"
	"testing"
)

func TestOutputToMap(t *testing.T) {
	assert.Any(
		t,
		map[string]string{"hello": "world"},
		show.OutputToMap("hello=world"),
	)
}
