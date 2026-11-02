package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/unit/client_tester"
	"runtime"
	"testing"
)

func TestTheClientReadsTheHost(t *testing.T) {
	c := client_tester.New(t, memory.New())
	h, e := c.GetHost(context.Background())
	assert.FatalOnError(t, e)
	assert.String(t, system.Hostname(), h.Hostname)
	assert.String(t, runtime.GOOS, h.Platform)
	assert.FatalOnError(t, c.Reachable(context.Background()))
}

func TestTheClientPassesTheOriginFilter(t *testing.T) {
	v, e := client_tester.New(t, memory.New()).ListServices(
		context.Background(),
		"unknown",
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}

func TestAFailedListingIsAnsweredAndCaptured(t *testing.T) {
	r := memory.New()
	c := client_tester.New(t, r)
	t.Setenv(constant.PathEnvironment, t.TempDir())
	v, e := c.ListServices(context.Background(), "")
	assert.Error(t, e)
	assert.Count(t, 0, v)
	assert.Count(t, 1, r.Events())
}
