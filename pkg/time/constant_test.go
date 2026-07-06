package time

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "15:04", HourMinute)
	assert.String(t, "15:04:05", HourMinuteSecond)
}
