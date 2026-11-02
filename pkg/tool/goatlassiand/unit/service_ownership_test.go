package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/failing_jira"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonClosesBothClientsItBuilt(t *testing.T) {
	j := service_tester.NewJira(t)
	c := service_tester.NewConfluence(t)
	d := service_tester.NewDaemon(
		t,
		service_tester.NewOption(j, c, nil),
		memory.New(),
	)
	d.Close()
	d.Close()
	assert.Integer(t, 1, j.Closed)
	assert.Integer(t, 1, c.Closed)
}

func TestTheDaemonHandsOutItsGuardsThenItsClock(t *testing.T) {
	workers := service_tester.NewDaemon(
		t,
		service_tester.NewOption(
			service_tester.NewJira(t),
			service_tester.NewConfluence(t),
			nil,
		),
		memory.New(),
	).Workers()
	assert.Count(t, 3, workers)

	for n, w := range workers {
		_, guard := w.(*peer.Guard)
		assert.Boolean(t, n < 2, guard)
	}
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	o := service_tester.NewOption(
		service_tester.NewJira(t),
		service_tester.NewConfluence(t),
		nil,
	)
	before := goleak.IgnoreCurrent()
	d := service_tester.NewDaemon(t, o, memory.New())
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}

func TestARefreshFillsThePlateAndNotifiesOnlyOnChange(t *testing.T) {
	j := service_tester.NewJira(t)
	j.Issues = []*issue.Issue{
		issue.New(issue.Raw("ALFA-1"), issue.FixtureOption()),
	}
	n := notifier.New()
	changes := n.Subscribe()
	defer n.Unsubscribe(changes)
	v := service_tester.NewService(
		t,
		j,
		service_tester.NewConfluence(t),
		nil,
		n,
	)
	assert.Count(t, 0, v.RefreshPlate())
	assert.Count(t, 1, v.CurrentPlate().Issues)
	assert.String(t, "ALFA-1", v.CurrentPlate().Issues[0].Key)
	assert.Integer(t, 1, len(changes))
	<-changes
	assert.Count(t, 0, v.RefreshPlate())
	assert.Integer(t, 0, len(changes))
}

func TestAFailedRefreshIsReportableAndKeepsThePlate(t *testing.T) {
	v := service_tester.NewService(
		t,
		failing_jira.New(errors.Format("search refused", "ALFA")),
		service_tester.NewConfluence(t),
		nil,
		notifier.New(),
	)
	failures := v.RefreshPlate()
	assert.Count(t, 1, failures)
	assert.True(t, classify.Reportable(failures[0]))
	assert.Count(t, 0, v.CurrentPlate().Issues)
}

func TestAnOutageIsLeftToTheGuard(t *testing.T) {
	v := service_tester.NewService(
		t,
		failing_jira.New(refusedConnection()),
		service_tester.NewConfluence(t),
		nil,
		notifier.New(),
	)
	failures := v.RefreshPlate()
	assert.Count(t, 1, failures)
	assert.False(t, classify.Reportable(failures[0]))
}
