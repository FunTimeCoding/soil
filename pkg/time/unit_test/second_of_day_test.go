package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/second"
	"testing"
	"time"
)

func TestOfDay(t *testing.T) {
	now := time.Now()
	y, m, d := now.Date()
	midnight := time.Date(
		y,
		m,
		d,
		0,
		0,
		0,
		0,
		now.Location(),
	)
	one := time.Date(
		y,
		m,
		d,
		1,
		0,
		0,
		0,
		now.Location(),
	)
	assert.Integer(t, 0, second.OfDay(midnight))
	assert.Integer(t, 3600, second.OfDay(one))
}
