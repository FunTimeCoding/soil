package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/backrest/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/worker"
	"testing"
)

func TestCommandOutageStaysQuietAndStopsAsking(t *testing.T) {
	c := mock_client.New()
	c.Failure = unavailable.Format("pgbackrest command unavailable: exit 125")
	r := memory.New()
	w, okay := daemon_tester.New(t, c, metric.New(), r).Workers()[1].(*worker.Worker)
	assert.True(t, okay)

	for range 3 {
		w.Poll()
	}

	assert.Count(t, 0, r.Events())
	assert.Integer(t, 1, c.Calls)
}

func TestPgbackrestFaultIsCapturedEachPoll(t *testing.T) {
	c := mock_client.New()
	c.Failure = errors.New("pgbackrest info: exit status 31")
	r := memory.New()
	w, okay := daemon_tester.New(t, c, metric.New(), r).Workers()[1].(*worker.Worker)
	assert.True(t, okay)
	w.Poll()
	w.Poll()
	assert.Count(t, 2, r.Events())
}
