package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/ranges"
	"testing"
	"time"
)

func TestLastWeek(t *testing.T) {
	assert.Any(t, time.Minute, ranges.LastWeek(1).Step)
}
