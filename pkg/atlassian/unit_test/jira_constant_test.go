package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"testing"
)

func TestJiraConstant(t *testing.T) {
	assert.Integer(t, 100, constant.SearchLimit)
	assert.String(t, "Assignee", constant.AssigneeName)
	assert.String(t, "Attachment", constant.AttachmentName)
	assert.String(t, "Canceled", constant.ServiceDeskCanceled)
	assert.String(t, "Development", constant.DevelopmentName)
	assert.String(t, "Flagged", constant.FlaggedName)
	assert.String(t, "Labels", constant.LabelsName)
	assert.String(t, "Linked Issues", constant.LinkedIssuesName)
	assert.String(t, "Parent", constant.ParentName)
	assert.String(t, "Rank", constant.RankName)
	assert.String(t, "Reporter", constant.ReporterName)
	assert.String(t, "Team", constant.TeamName)
}
