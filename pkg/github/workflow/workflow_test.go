package workflow

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v88/github"
	"testing"
	"time"
)

func TestWorkflow(t *testing.T) {
	r := New(
		&github.Workflow{
			Name:      new(upper.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &Workflow{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
