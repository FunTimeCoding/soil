package booleans

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestToString(t *testing.T) {
	assert.String(t, "1", ToString(true))
	assert.String(t, "0", ToString(false))
}
