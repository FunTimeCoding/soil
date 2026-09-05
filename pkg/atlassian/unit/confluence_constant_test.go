package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"testing"
)

func TestClient(t *testing.T) {
	assert.String(t, "no space", constant.ConfluenceNoSpace)
	assert.String(t, "view", constant.ConfluenceViewFormat)
	assert.String(t, "atlas_doc_format", constant.ConfluenceAtlasFormat)
	assert.String(t, "export_view", constant.ConfluenceExportFormat)
	assert.String(
		t,
		"anonymous_export_view",
		constant.ConfluenceAnonymousFormat,
	)
	assert.String(t, "styled_view", constant.ConfluenceStyledFormat)
	assert.String(t, "editor", constant.ConfluenceEditFormat)
}
