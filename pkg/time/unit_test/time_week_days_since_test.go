package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/time"
	"testing"
	"time"
)

func TestWorkDaysSince(t *testing.T) {
	assert.Integer(
		t,
		3,
		library.WeekDaysSince(
			library.NewDay(2024, time.October, 25),
			library.NewDay(2024, time.October, 30),
		),
	)
}
