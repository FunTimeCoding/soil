package job

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
	"time"
)

func TestJob(t *testing.T) {
	r := New(
		&github.WorkflowJob{
			Name:      new(upper.Alfa),
			CreatedAt: &github.Timestamp{},
		},
	)
	r.Raw = nil
	assert.Any(t, &Job{Name: "Alfa", CreatedAt: time.Time{}}, r)
}
