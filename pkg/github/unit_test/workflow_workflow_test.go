package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/workflow"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
	"time"
)

func TestWorkflow(t *testing.T) {
	r := workflow.New(
		&github.Workflow{
			Name:      new(upper.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &workflow.Workflow{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
