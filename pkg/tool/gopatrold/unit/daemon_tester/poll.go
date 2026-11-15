package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/worker"
	"testing"
)

func Poll(
	t *testing.T,
	d *daemon.Daemon,
) {
	t.Helper()
	w := d.Workers()
	clock, okay := w[len(w)-1].(*worker.Worker)
	assert.True(t, okay)
	clock.Poll()
}
