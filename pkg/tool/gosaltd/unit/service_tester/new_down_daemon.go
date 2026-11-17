package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/salt/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/daemon"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
	"time"
)

func NewDownDaemon(
	t *testing.T,
	r *memory.Memory,
) *daemon.Daemon {
	t.Helper()
	c := mock_client.New(t)
	c.Fail(c.Reachable, outage_tester.RefusedDial())
	c.Fail(c.Keys, outage_tester.RefusedDial())
	c.Fail(c.Minions, outage_tester.RefusedDial())

	return New(
		t,
		WithReporter(r),
		WithMaster(func() face.SaltClient { return c }),
		WithKeySyncInterval(20*time.Millisecond),
	).Daemon()
}
