package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/unit/client_tester"
	"runtime"
	"testing"
)

func TestTheClientReadsTheHost(t *testing.T) {
	c := client_tester.New(t)
	h, e := c.Host(context.Background())
	assert.FatalOnError(t, e)
	assert.String(t, system.Hostname(), h.Hostname)
	assert.String(t, runtime.GOOS, h.Platform)
	assert.FatalOnError(t, c.Reachable(context.Background()))
}

func TestTheClientPassesTheOriginFilter(t *testing.T) {
	v, e := client_tester.New(t).Services(context.Background(), "unknown")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}
