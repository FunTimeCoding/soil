package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/text/markdown"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"[Alfa](Bravo)",
		markdown.Link(upper.Alfa, upper.Bravo),
	)
}
