package unit_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"testing"
)

func TestNotConfiguredClassifiesByType(t *testing.T) {
	e := not_configured.Format("no default channel configured")
	assert.String(t, "no default channel configured", e.Error())
	assert.True(t, not_configured.Is(e))
	assert.False(t, not_configured.Is(errors.New("no default channel")))
}

func TestUnreachableClassifiesByType(t *testing.T) {
	e := unreachable.Format("extension not connected")
	assert.String(t, "extension not connected", e.Error())
	assert.True(t, unreachable.Is(e))
	assert.False(t, unreachable.Is(errors.New("extension not connected")))
}

func TestTimeoutClassifiesByType(t *testing.T) {
	e := timeout.Format("job timeout: %s", "sweep")
	assert.String(t, "job timeout: sweep", e.Error())
	assert.True(t, timeout.Is(e))
	assert.True(t, timeout.Is(fmt.Errorf("wait: %w", context.DeadlineExceeded)))
	assert.False(t, timeout.Is(errors.New("job timeout: sweep")))
}

func TestUnexpectedClassifiesByType(t *testing.T) {
	e := unexpected.Format("unexpected status: %d", 502)
	assert.String(t, "unexpected status: 502", e.Error())
	assert.True(t, unexpected.Is(e))
	assert.False(t, unexpected.Is(errors.New("unexpected status: 502")))
}
