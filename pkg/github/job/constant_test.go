package job

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "completed", Completed)
	assert.String(t, "in_progress", InProgress)
	assert.String(t, "queued", Queued)
	assert.String(t, "success", Success)
}
