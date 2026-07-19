package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/ranges"
	"testing"
	"time"
)

func TestLastDay(t *testing.T) {
	assert.Any(t, time.Minute, ranges.LastDay(1).Step)
}
