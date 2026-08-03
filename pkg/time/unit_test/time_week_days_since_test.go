package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/time/day"
	"testing"
	"time"
)

func TestWorkDaysSince(t *testing.T) {
	assert.Integer(
		t,
		3,
		library.WeekDaysSince(
			day.New(2024, time.October, 25),
			day.New(2024, time.October, 30),
		),
	)
}
