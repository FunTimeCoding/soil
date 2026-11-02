package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/text/markdown"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"[Alfa](Bravo)",
		markdown.Link(stringsConstant.UpperAlfa, stringsConstant.UpperBravo),
	)
}
