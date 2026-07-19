package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"testing"
)

func TestJiraIssueConstant(t *testing.T) {
	assert.String(t, "Bug", issue.BugType)
	assert.String(t, "Story", issue.StoryType)
	assert.String(t, "Task", issue.TaskType)
	assert.String(t, "Sub-task", issue.SubTaskType)
}
