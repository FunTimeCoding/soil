package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/mock_source"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/types/run_key"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/unit/service_tester"
	"testing"
)

func TestLatestKeepsTheNewestCompletedRun(t *testing.T) {
	s := mock_source.New()
	s.AddRun("bravo", workflowRun(1, "main", "completed", "failure"))
	s.AddRun("bravo", workflowRun(3, "main", "completed", "success"))
	s.AddRun("bravo", workflowRun(2, "main", "completed", "failure"))
	v, e := service_tester.New(t, s).Latest()
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, len(v))
	assert.String(
		t,
		"success",
		v[run_key.Key{
			Repository: "bravo",
			Workflow:   "release image",
			Branch:     "main",
		}].Conclusion,
	)
}

func TestLatestSkipsARunInProgress(t *testing.T) {
	s := mock_source.New()
	s.AddRun("bravo", workflowRun(1, "main", "completed", "success"))
	s.AddRun("bravo", workflowRun(2, "main", "in_progress", ""))
	v, e := service_tester.New(t, s).Latest()
	assert.FatalOnError(t, e)
	assert.Integer(
		t,
		int64(1),
		v[run_key.Key{
			Repository: "bravo",
			Workflow:   "release image",
			Branch:     "main",
		}].Identifier,
	)
}

func TestLatestKeepsBranchesApart(t *testing.T) {
	s := mock_source.New()
	s.AddRun("bravo", workflowRun(1, "main", "completed", "success"))
	s.AddRun("bravo", workflowRun(2, "feature", "completed", "failure"))
	s.AddRun("charlie", workflowRun(3, "main", "completed", "success"))
	v, e := service_tester.New(t, s).Latest()
	assert.FatalOnError(t, e)
	assert.Integer(t, 3, len(v))
}
