package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/expression"
	"testing"
)

func TestSubMatch(t *testing.T) {
	assert.String(
		t,
		"123",
		expression.SubMatch(`MyPrefix: (.*)`, "MyPrefix: 123"),
	)
}
