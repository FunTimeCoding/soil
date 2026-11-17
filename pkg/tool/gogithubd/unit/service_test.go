package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/types/run_key"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/unit/service_tester"
	"testing"
)

func TestLatestKeepsNewestCompletedRun(t *testing.T) {
	o := service_tester.New(t)
	o.Source.AddRun(
		"bravo",
		service_tester.WorkflowRun(1, "main", "completed", "failure"),
	)
	o.Source.AddRun(
		"bravo",
		service_tester.WorkflowRun(3, "main", "completed", "success"),
	)
	o.Source.AddRun(
		"bravo",
		service_tester.WorkflowRun(2, "main", "completed", "failure"),
	)
	v, e := o.Service().Latest()
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

func TestLatestSkipsRunInProgress(t *testing.T) {
	o := service_tester.New(t)
	o.Source.AddRun(
		"bravo",
		service_tester.WorkflowRun(1, "main", "completed", "success"),
	)
	o.Source.AddRun(
		"bravo",
		service_tester.WorkflowRun(2, "main", "in_progress", ""),
	)
	v, e := o.Service().Latest()
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
	o := service_tester.New(t)
	o.Source.AddRun(
		"bravo",
		service_tester.WorkflowRun(1, "main", "completed", "success"),
	)
	o.Source.AddRun(
		"bravo",
		service_tester.WorkflowRun(2, "feature", "completed", "failure"),
	)
	o.Source.AddRun(
		"charlie",
		service_tester.WorkflowRun(3, "main", "completed", "success"),
	)
	v, e := o.Service().Latest()
	assert.FatalOnError(t, e)
	assert.Integer(t, 3, len(v))
}
