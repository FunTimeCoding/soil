package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"testing"
)

func TestJiraIssueConstant(t *testing.T) {
	assert.String(t, "Bug", constant.JiraBugType)
	assert.String(t, "Story", constant.JiraStoryType)
	assert.String(t, "Task", constant.JiraTaskType)
	assert.String(t, "Sub-task", constant.JiraSubTaskType)
}
