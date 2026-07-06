package issue

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "Bug", BugType)
	assert.String(t, "Story", StoryType)
	assert.String(t, "Task", TaskType)
	assert.String(t, "Sub-task", SubTaskType)
}
