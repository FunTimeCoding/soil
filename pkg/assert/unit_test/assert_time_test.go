package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	assert.Time(t, now, now)
}
