package second

import (
	"github.com/funtimecoding/soil/pkg/assert"
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
	assert.Integer(t, 0, OfDay(midnight))
	assert.Integer(t, 3600, OfDay(one))
}
