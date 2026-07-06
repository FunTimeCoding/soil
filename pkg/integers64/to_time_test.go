package integers64

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time"
	"testing"
)

func TestToTime(t *testing.T) {
	assert.String(
		t,
		"1970-01-01 00:00",
		ToTime(0).Format(time.DateMinute),
	)
}
