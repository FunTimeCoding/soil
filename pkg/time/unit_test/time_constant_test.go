package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "15:04", time.HourMinute)
	assert.String(t, "15:04:05", time.HourMinuteSecond)
}
