package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/time"
	"testing"
	"time"
)

func TestOlderThan(t *testing.T) {
	now := time.Now()
	assert.False(t, library.OlderThan(now, 1))
	fiveMinutesAgo := now.Add(-5 * time.Minute)
	assert.True(t, library.OlderThan(fiveMinutesAgo, 1))
}
