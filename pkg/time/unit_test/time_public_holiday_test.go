package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/time"
	"testing"
	"time"
)

func TestPublicHoliday(t *testing.T) {
	assert.True(
		t,
		library.PublicHoliday(library.NewDay(2024, time.November, 1)),
	)
	assert.False(
		t,
		library.PublicHoliday(library.NewDay(2024, time.November, 2)),
	)
}
