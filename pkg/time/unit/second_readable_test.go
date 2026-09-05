package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/time/second"
	"testing"
)

func TestSecondReadable(t *testing.T) {
	minute := constant.MinuteInSeconds
	hour := constant.HourInSeconds
	day := constant.DayInSeconds
	month := constant.MonthInSeconds
	assert.String(t, "1 second", second.Readable(1))
	assert.String(t, "59 seconds", second.Readable(59))
	assert.String(t, "1 minute", second.Readable(minute))
	assert.String(t, "30 minutes", second.Readable(minute*30))
	assert.String(t, "59 minutes", second.Readable(minute*59))
	assert.String(t, "1 hour", second.Readable(hour))
	assert.String(t, "1 day", second.Readable(day))
	assert.String(t, "1.5 days", second.Readable(hour*36))
	assert.String(t, "2 days", second.Readable(day*2))
	assert.String(t, "1 week", second.Readable(day*7))
	assert.String(t, "1 month", second.Readable(month))
}
