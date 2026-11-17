package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/worker_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestAlertmanagerOutageIsLeftToGuard(t *testing.T) {
	o := worker_tester.New(t)
	o.Alertmanager.Fail(outage_tester.RefusedDial())
	o.Worker.Poll()
	o.Worker.Poll()
	assert.Count(t, 0, o.Reporter.Events())
	_, e := o.Service.Record()
	assert.True(t, unavailable.Is(e))
}

func TestAlertmanagerFaultIsCapturedOnce(t *testing.T) {
	o := worker_tester.New(t)
	o.Alertmanager.Fail(errors.Format("alertmanager answered", "garbage"))
	o.Worker.Poll()
	assert.Count(t, 1, o.Reporter.Events())
}
