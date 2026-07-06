package run

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v88/github"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	r := New(
		&github.WorkflowRun{
			Name:       new(upper.Alfa),
			CreatedAt:  &github.Timestamp{},
			Repository: &github.Repository{},
		},
	)
	r.Repository = nil
	r.Raw = nil
	assert.Any(
		t,
		&Run{
			MonitorIdentifier: "ghjob-0",
			Name:              "Alfa",
			Create:            time.Time{},
		},
		r,
	)
}
