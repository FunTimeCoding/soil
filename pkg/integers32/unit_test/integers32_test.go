package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers32"
	"testing"
)

func TestToString(t *testing.T) {
	assert.String(t, "1", integers32.ToString(1))
}
