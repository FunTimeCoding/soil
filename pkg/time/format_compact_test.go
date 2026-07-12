package time

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestFormatCompactToday(t *testing.T) {
	now := time.Now()
	assert.String(t, now.Local().Format(HourMinute), FormatCompact(now))
}

func TestFormatCompactOlder(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1)
	assert.String(
		t,
		yesterday.Local().Format(DateMinute),
		FormatCompact(yesterday),
	)
}
