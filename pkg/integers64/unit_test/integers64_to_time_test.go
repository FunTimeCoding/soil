package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers64"
	"github.com/funtimecoding/soil/pkg/time"
	"testing"
)

func TestToTime(t *testing.T) {
	assert.String(
		t,
		"1970-01-01 00:00",
		integers64.ToTime(0).Format(time.DateMinute),
	)
}
