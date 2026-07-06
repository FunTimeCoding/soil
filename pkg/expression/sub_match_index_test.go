package expression

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSubMatchIndex(t *testing.T) {
	assert.String(
		t,
		"MyPrefix: 123",
		SubMatchIndex(
			`MyPrefix: (.*)`,
			"Some line\nMyPrefix: 123\nSome other line\n",
			0,
		),
	)
}
