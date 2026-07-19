package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/second"
	"testing"
)

func TestToClock(t *testing.T) {
	assert.String(t, "00:00", second.ToClock(0))
}
