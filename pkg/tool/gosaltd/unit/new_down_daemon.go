package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/salt/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/daemon"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
	"time"
)

func newDownDaemon(
	t *testing.T,
	r *memory.Memory,
) *daemon.Daemon {
	t.Helper()
	c := mock_client.New()
	c.Fail(c.Reachable, refusedDial())
	c.Fail(c.Keys, refusedDial())
	c.Fail(c.Minions, refusedDial())
	o := newOption(t)
	o.KeySyncInterval = 20 * time.Millisecond
	o.Master = func() face.SaltClient { return c }
	result := daemon.New(
		o,
		logger.New(t.Context()),
		r,
		prometheus.NewRegistry(),
	)
	t.Cleanup(result.Close)

	return result
}
