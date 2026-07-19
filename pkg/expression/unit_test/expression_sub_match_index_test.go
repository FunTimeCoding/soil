package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/expression"
	"testing"
)

func TestSubMatchIndex(t *testing.T) {
	assert.String(
		t,
		"MyPrefix: 123",
		expression.SubMatchIndex(
			`MyPrefix: (.*)`,
			"Some line\nMyPrefix: 123\nSome other line\n",
			0,
		),
	)
}
