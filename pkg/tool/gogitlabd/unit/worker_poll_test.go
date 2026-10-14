package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"testing"
)

func TestPollKeepsTheRowsOfAFailingInstance(t *testing.T) {
	alfa := seededForge(1, "app")
	bravo := seededForge(2, "web")
	r := memory.New()
	w := newWorker(t, alfa, bravo, r)
	defer w.Close()
	w.Poll()
	assert.Count(t, 2, w.Entries())
	alfa.FailProjects(errors.New("unreachable"))
	bravo.SeedBranch("feature", "bbb")
	bravo.SeedPipeline(11, "feature", "bbb", "running")
	w.Poll()
	var instances []string

	for _, e := range w.Entries() {
		instances = append(instances, e.Instance)
	}

	assert.Strings(t, []string{"alfa", "bravo", "bravo"}, instances)
	assert.Count(t, 1, r.Events())
}

func TestRequestsCarryTheirInstance(t *testing.T) {
	alfa := seededForge(1, "app")
	alfa.SetMergeRequests(
		[]*merge_request.Request{newRequest(1, 1, "alfa change", 1)},
		nil,
	)
	bravo := seededForge(2, "web")
	bravo.SetMergeRequests(
		nil,
		[]*merge_request.Request{newRequest(2, 1, "bravo change", 2)},
	)
	w := newWorker(t, alfa, bravo, memory.New())
	defer w.Close()
	w.Poll()
	requests := w.Requests()
	assert.Count(t, 2, requests)
	assert.String(t, "alfa", requests[0].Instance)
	assert.String(t, "alfa change", requests[0].Request.Title)
	assert.String(t, "bravo", requests[1].Instance)
}
