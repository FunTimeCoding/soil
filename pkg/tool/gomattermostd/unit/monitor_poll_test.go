package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"testing"
)

func TestThePollReportsAFault(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	r := memory.New()
	w := worker_tester.NewMonitor(t, m, r, "alfa")
	m.Failure = fault()
	w.Poll()
	assert.Count(t, 1, r.Events())
}

func TestThePollLeavesAnOutageToTheGuard(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	r := memory.New()
	w := worker_tester.NewMonitor(t, m, r, "alfa")
	m.Failure = unavailable.Format("mattermost maintenance")
	w.Poll()
	assert.Count(t, 0, r.Events())
}
