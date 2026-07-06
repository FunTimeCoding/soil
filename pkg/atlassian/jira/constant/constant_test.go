package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Integer(t, 100, SearchLimit)
	assert.String(t, "Assignee", AssigneeName)
	assert.String(t, "Attachment", AttachmentName)
	assert.String(t, "Canceled", ServiceDeskCanceled)
	assert.String(t, "Development", DevelopmentName)
	assert.String(t, "Flagged", FlaggedName)
	assert.String(t, "Labels", LabelsName)
	assert.String(t, "Linked Issues", LinkedIssuesName)
	assert.String(t, "Parent", ParentName)
	assert.String(t, "Rank", RankName)
	assert.String(t, "Reporter", ReporterName)
	assert.String(t, "Team", TeamName)
}
