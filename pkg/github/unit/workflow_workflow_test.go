package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/workflow"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/google/go-github/v90/github"
	"testing"
	"time"
)

func TestWorkflow(t *testing.T) {
	r := workflow.New(
		&github.Workflow{
			Name:      new(constant.UpperAlfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &workflow.Workflow{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
