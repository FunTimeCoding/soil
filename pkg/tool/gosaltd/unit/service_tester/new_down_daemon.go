package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/salt/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/daemon"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
	"time"
)

func NewDownDaemon(
	t *testing.T,
	r *memory.Memory,
) *daemon.Daemon {
	t.Helper()
	c := mock_client.New()
	c.Fail(c.Reachable, outage_tester.RefusedDial())
	c.Fail(c.Keys, outage_tester.RefusedDial())
	c.Fail(c.Minions, outage_tester.RefusedDial())
	o := NewOption(t)
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
