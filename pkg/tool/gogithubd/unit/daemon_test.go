package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/face"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/mock_source"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/unit/closing_source"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/unit/daemon_tester"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"go.uber.org/goleak"
	"strings"
	"testing"
)

func TestTheDaemonClosesTheSourceItBuilt(t *testing.T) {
	s := closing_source.New()
	o := option.New()
	o.Source = func() face.RunSource { return s }
	daemon.New(o, logger.New(t.Context()), memory.New(), metric.New()).Close()
	assert.Integer(t, 1, s.Closed())
}

func TestTheClockGaugesTheOwnersLatestRuns(t *testing.T) {
	s := mock_source.New()
	s.AddRun("bravo", workflowRun(1, "main", "completed", "failure"))
	s.AddRun("bravo", workflowRun(2, "main", "completed", "success"))
	m := metric.New()

	for _, w := range daemon_tester.New(t, s, m).Workers() {
		w.Start()
		w.Close()
	}

	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			m.Registry(),
			strings.NewReader(
				`
# HELP github_workflow_run_status Latest completed workflow run per repo/workflow/branch (1 per entry)
# TYPE github_workflow_run_status gauge
github_workflow_run_status{branch="main",conclusion="success",owner="alfa",repo="bravo",workflow="release image"} 1
`,
			),
			"github_workflow_run_status",
		),
	)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	s := mock_source.New()
	o := option.New()
	o.Source = func() face.RunSource { return s }
	d := daemon.New(o, logger.New(t.Context()), memory.New(), metric.New())
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
