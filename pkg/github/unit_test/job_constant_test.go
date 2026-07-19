package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/job"
	"testing"
)

func TestJobConstant(t *testing.T) {
	assert.String(t, "completed", job.Completed)
	assert.String(t, "in_progress", job.InProgress)
	assert.String(t, "queued", job.Queued)
	assert.String(t, "success", job.Success)
}
