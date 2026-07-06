package console

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestLink(t *testing.T) {
	assert.Any(
		t,
		"\x1b]8;;https://example.org\x1b\\Example\x1b]8;;\x1b\\",
		Link("https://example.org", "Example", true),
	)
}
