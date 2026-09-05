package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"testing"
)

func TestJiraConstant(t *testing.T) {
	assert.Integer(t, 100, constant.JiraSearchLimit)
	assert.String(t, "Assignee", constant.JiraAssigneeName)
	assert.String(t, "Attachment", constant.JiraAttachmentName)
	assert.String(t, "Canceled", constant.ServiceDeskCanceled)
	assert.String(t, "Development", constant.JiraDevelopmentName)
	assert.String(t, "Flagged", constant.JiraFlaggedName)
	assert.String(t, "Labels", constant.JiraLabelsName)
	assert.String(t, "Linked Issues", constant.JiraLinkedIssuesName)
	assert.String(t, "Parent", constant.JiraParentName)
	assert.String(t, "Rank", constant.JiraRankName)
	assert.String(t, "Reporter", constant.JiraReporterName)
	assert.String(t, "Team", constant.JiraTeamName)
}
