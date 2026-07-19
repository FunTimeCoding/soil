package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text"
	"testing"
)

func TestPercent(t *testing.T) {
	assert.String(t, "A (1%)", text.FormatPercent("A", 1.2))
}
