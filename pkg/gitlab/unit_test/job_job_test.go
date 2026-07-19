package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestJob(t *testing.T) {
	a := job.New(&gitlab.Job{})
	a.Raw = nil
	assert.Any(t, &job.Job{MonitorIdentifier: "gitlab-0"}, a)
}
