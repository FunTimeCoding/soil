package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"testing"
)

func TestStatusConstant(t *testing.T) {
	assert.String(t, "completed", constant.CompletedStatus)
	assert.String(t, "in_progress", constant.InProgressStatus)
	assert.String(t, "queued", constant.QueuedStatus)
	assert.String(t, "success", constant.SuccessConclusion)
	assert.String(t, "failure", constant.FailureConclusion)
	assert.String(t, "failed", constant.RunFailedConcern)
}
