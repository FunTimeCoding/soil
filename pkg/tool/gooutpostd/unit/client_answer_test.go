package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/unit/client_tester"
	"runtime"
	"testing"
)

func TestClientReadsHost(t *testing.T) {
	c := client_tester.New(t, base.New(t))
	h, e := c.GetHost(context.Background())
	assert.FatalOnError(t, e)
	assert.String(t, system.Hostname(), h.Hostname)
	assert.String(t, runtime.GOOS, h.Platform)
	assert.FatalOnError(t, c.Reachable(context.Background()))
}

func TestClientPassesOriginFilter(t *testing.T) {
	v, e := client_tester.New(t, base.New(t)).ListServices(
		context.Background(),
		"unknown",
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}

func TestFailedListingIsAnsweredAndCaptured(t *testing.T) {
	s := base.New(t)
	c := client_tester.New(t, s)
	t.Setenv(systemConstant.PathEnvironment, t.TempDir())
	v, e := c.ListServices(context.Background(), "")
	assert.Error(t, e)
	assert.Count(t, 0, v)
	assert.Count(t, 1, s.Reporter.Events())
}
