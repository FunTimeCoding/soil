package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers64"
	"testing"
)

func TestToString(t *testing.T) {
	assert.String(t, "1", integers64.ToString(1))
}
