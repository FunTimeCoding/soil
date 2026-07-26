package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "15:04", constant.HourMinute)
	assert.String(t, "15:04:05", constant.HourMinuteSecond)
}
