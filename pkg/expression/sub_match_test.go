package expression

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSubMatch(t *testing.T) {
	assert.String(
		t,
		"123",
		SubMatch(`MyPrefix: (.*)`, "MyPrefix: 123"),
	)
}
