package markdown

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"[Alfa](Bravo)",
		Link(upper.Alfa, upper.Bravo),
	)
}
