package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/status"
	"testing"
)

func TestStatusActiveCoversEveryUnfinishedState(t *testing.T) {
	assert.True(t, status.Active(constant.JobRunning))
	assert.True(t, status.Active(constant.JobPending))
	assert.True(t, status.Active(constant.JobCreated))
	assert.True(t, status.Active(constant.JobPreparing))
	assert.True(t, status.Active(constant.JobWaitingForResource))
}

func TestStatusActiveRejectsFinishedStates(t *testing.T) {
	assert.False(t, status.Active(constant.JobSuccess))
	assert.False(t, status.Active(constant.JobFail))
	assert.False(t, status.Active(constant.JobCanceled))
	assert.False(t, status.Active(constant.JobSkipped))
	assert.False(t, status.Active(constant.JobManual))
	assert.False(t, status.Active(""))
}
