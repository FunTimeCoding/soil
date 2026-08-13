package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/time"
	"testing"
	"time"
)

func TestFormatCompactToday(t *testing.T) {
	now := time.Now()
	assert.String(t, now.Local().Format("15:04"), library.FormatCompact(now))
}

func TestFormatCompactOlder(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1)
	assert.String(
		t,
		yesterday.Local().Format("2006-01-02 15:04"),
		library.FormatCompact(yesterday),
	)
}
