package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/types/run_key"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"strings"
	"testing"
	"time"
)

func TestPollDropsRunNoLongerLatest(t *testing.T) {
	branches := []string{"main", "feature"}
	y := prometheus.NewRegistry()
	w := worker.New(
		func() (map[run_key.Key]*run.Run, error) {
			b := branches[0]
			branches = branches[1:]

			return map[run_key.Key]*run.Run{
				{
					Repository: "bravo",
					Workflow:   "release image",
					Branch:     b,
				}: workflowRun(1, b, "completed", "success"),
			}, nil
		},
		"alfa",
		time.Minute,
		logger.New(t.Context()),
		y,
		memory.New(),
	)
	defer w.Close()
	w.Poll()
	w.Poll()
	assert.FatalOnError(
		t,
		testutil.GatherAndCompare(
			y,
			strings.NewReader(
				`
# HELP github_workflow_run_status Latest completed workflow run per repo/workflow/branch (1 per entry)
# TYPE github_workflow_run_status gauge
github_workflow_run_status{branch="feature",conclusion="success",owner="alfa",repo="bravo",workflow="release image"} 1
`,
			),
			"github_workflow_run_status",
		),
	)
}
