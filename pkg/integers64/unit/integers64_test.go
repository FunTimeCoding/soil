package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers64"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"testing"
)

func TestTo32(t *testing.T) {
	assert.Integer(t, 0, integers64.To32(0))
}

func TestToString(t *testing.T) {
	assert.String(t, "1", integers64.ToString(1))
}

func TestToTime(t *testing.T) {
	assert.String(
		t,
		"1970-01-01 00:00",
		integers64.ToTime(0).Format(constant.DateMinute),
	)
}

func TestToUnsigned32(t *testing.T) {
	assert.Integer(t, 0, integers64.ToUnsigned32(0))
}
