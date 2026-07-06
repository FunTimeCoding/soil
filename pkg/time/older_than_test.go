package time

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestOlderThan(t *testing.T) {
	now := time.Now()
	assert.False(t, OlderThan(now, 1))
	fiveMinutesAgo := now.Add(-5 * time.Minute)
	assert.True(t, OlderThan(fiveMinutesAgo, 1))
}
