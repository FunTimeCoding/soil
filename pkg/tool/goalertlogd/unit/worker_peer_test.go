package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/worker_tester"
	"testing"
	"time"
)

func TestAnAlertmanagerOutageIsLeftToTheGuard(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Fail(refusedConnection())
	o.Worker.Poll()
	o.Worker.Poll()
	assert.Count(t, 0, o.Reporter.Events())
	_, e := o.Service.Record(time.Now())
	assert.True(t, unavailable.Is(e))
}

func TestAnAlertmanagerFaultIsCapturedOnce(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Fail(errors.Format("alertmanager answered", "garbage"))
	o.Worker.Poll()
	assert.Count(t, 1, o.Reporter.Events())
}

func TestTheGuardRunsOnTheLifecycle(t *testing.T) {
	o := worker_tester.New(t)
	assert.Count(t, 1, o.Service.Workers())
}
