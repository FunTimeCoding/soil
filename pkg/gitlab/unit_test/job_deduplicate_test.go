package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	assert.Count(t, 0, job.Deduplicate([]*gitlab.Job{}))
}
