package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/job"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
	"time"
)

func TestJob(t *testing.T) {
	r := job.New(
		&github.WorkflowJob{
			Name:      new(upper.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &job.Job{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
