package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/time"
	"testing"
	"time"
)

func TestRelative(t *testing.T) {
	now := time.Now()
	assert.String(t, "just now", library.Relative(now.Add(-time.Second)))
	assert.String(t, "5m ago", library.Relative(now.Add(-5*time.Minute)))
	assert.String(t, "3h ago", library.Relative(now.Add(-3*time.Hour)))
	assert.String(t, "2d ago", library.Relative(now.Add(-50*time.Hour)))
}
