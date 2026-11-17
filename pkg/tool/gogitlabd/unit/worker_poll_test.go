package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/gitlab_tester"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/worker_tester"
	"testing"
)

func TestPollReportsRefusalBesideInstanceThatIsAway(t *testing.T) {
	o := worker_tester.New(t)
	w := o.Worker()
	service_tester.SeedForge(o.Alfa(), 1, "app")
	service_tester.SeedForge(o.Bravo(), 2, "web")
	o.Alfa().FailProjects(unavailable.Format("gitlab away"))
	o.Bravo().FailProjects(errors.New("refused for no known reason"))
	w.Poll()
	events := o.Reporter.Events()
	assert.Count(t, 1, events)
	assert.String(t, "refused for no known reason", events[0].Error.Error())
}

func TestPollKeepsRowsOfFailingInstance(t *testing.T) {
	o := worker_tester.New(t)
	w := o.Worker()
	s := o.Service()
	alfa := o.Alfa()
	bravo := o.Bravo()
	service_tester.SeedForge(alfa, 1, "app")
	service_tester.SeedForge(bravo, 2, "web")
	w.Poll()
	assert.Count(t, 2, s.Entries())
	alfa.FailProjects(errors.New("unreachable"))
	bravo.SeedBranch("feature", "bbb")
	bravo.SeedPipeline(11, "feature", "bbb", "running")
	w.Poll()
	var instances []string

	for _, e := range s.Entries() {
		instances = append(instances, e.Instance)
	}

	assert.Strings(t, []string{"alfa", "bravo", "bravo"}, instances)
	assert.Count(t, 1, o.Reporter.Events())
}

func TestRequestsCarryTheirInstance(t *testing.T) {
	o := worker_tester.New(t)
	w := o.Worker()
	s := o.Service()
	alfa := o.Alfa()
	service_tester.SeedForge(alfa, 1, "app")
	alfa.SetMergeRequests(
		[]*merge_request.Request{gitlab_tester.Request(1, 1, "alfa change", 1)},
		nil,
	)
	bravo := o.Bravo()
	service_tester.SeedForge(bravo, 2, "web")
	bravo.SetMergeRequests(
		nil,
		[]*merge_request.Request{
			gitlab_tester.Request(2, 1, "bravo change", 2),
		},
	)
	w.Poll()
	requests := s.Requests()
	assert.Count(t, 2, requests)
	assert.String(t, "alfa", requests[0].Instance)
	assert.String(t, "alfa change", requests[0].Request.Title)
	assert.String(t, "bravo", requests[1].Instance)
}
