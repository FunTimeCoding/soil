package job

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestJob(t *testing.T) {
	a := New(&gitlab.Job{})
	a.Raw = nil
	assert.Any(t, &Job{MonitorIdentifier: "gitlab-0"}, a)
}
