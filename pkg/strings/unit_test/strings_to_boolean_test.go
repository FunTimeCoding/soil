package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestToBoolean(t *testing.T) {
	assert.True(t, strings.ToBoolean(strings.BooleanTrue))
	assert.False(t, strings.ToBoolean(strings.BooleanFalse))
}
