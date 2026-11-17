package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesBothClientsItBuilt(t *testing.T) {
	s := service_tester.New(t)
	d := s.Daemon()
	d.Close()
	d.Close()
	assert.Integer(t, 1, s.Jira.Closed)
	assert.Integer(t, 1, s.Confluence.Closed)
}

func TestDaemonHandsOutItsGuardsThenItsClock(t *testing.T) {
	workers := service_tester.New(t).Daemon().Workers()
	assert.Count(t, 3, workers)

	for n, w := range workers {
		_, guard := w.(*peer.Guard)
		assert.Boolean(t, n < 2, guard)
	}
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := service_tester.New(t)
	before := goleak.IgnoreCurrent()
	d := s.Daemon()
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

func TestRefreshFillsPlateAndNotifiesOnlyOnChange(t *testing.T) {
	n := notifier.New()
	changes := n.Subscribe()
	defer n.Unsubscribe(changes)
	s := service_tester.New(t, service_tester.WithNotifier(n))
	s.Jira.Issues = []*issue.Issue{
		issue.New(issue.Raw("ALFA-1"), issue.FixtureOption()),
	}
	v := s.Service()
	assert.Count(t, 0, v.RefreshPlate())
	assert.Count(t, 1, v.CurrentPlate().Issues)
	assert.String(t, "ALFA-1", v.CurrentPlate().Issues[0].Key)
	assert.Integer(t, 1, len(changes))
	<-changes
	assert.Count(t, 0, v.RefreshPlate())
	assert.Integer(t, 0, len(changes))
}

func TestFailedRefreshIsReportableAndKeepsPlate(t *testing.T) {
	s := service_tester.New(t)
	s.Jira.Fail(s.Jira.Search, errors.Format("search refused", "ALFA"))
	v := s.Service()
	failures := v.RefreshPlate()
	assert.Count(t, 1, failures)
	assert.True(t, classify.Reportable(failures[0]))
	assert.Count(t, 0, v.CurrentPlate().Issues)
}

func TestOutageIsLeftToGuard(t *testing.T) {
	s := service_tester.New(t)
	s.Jira.Fail(s.Jira.Search, outage_tester.RefusedDial())
	failures := s.Service().RefreshPlate()
	assert.Count(t, 1, failures)
	assert.False(t, classify.Reportable(failures[0]))
}
