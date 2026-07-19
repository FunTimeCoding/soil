package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console"
	"testing"
)

func TestLink(t *testing.T) {
	assert.Any(
		t,
		"\x1b]8;;https://example.org\x1b\\Example\x1b]8;;\x1b\\",
		console.Link("https://example.org", "Example", true),
	)
}
