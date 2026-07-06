package second

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time"
	"testing"
)

func TestReadable(t *testing.T) {
	minute := time.MinuteInSeconds
	hour := time.HourInSeconds
	day := time.DayInSeconds
	month := time.MonthInSeconds
	assert.String(t, "1 second", Readable(1))
	assert.String(t, "59 seconds", Readable(59))
	assert.String(t, "1 minute", Readable(minute))
	assert.String(t, "30 minutes", Readable(minute*30))
	assert.String(t, "59 minutes", Readable(minute*59))
	assert.String(t, "1 hour", Readable(hour))
	assert.String(t, "1 day", Readable(day))
	assert.String(t, "1.5 days", Readable(hour*36))
	assert.String(t, "2 days", Readable(day*2))
	assert.String(t, "1 week", Readable(day*7))
	assert.String(t, "1 month", Readable(month))
}
