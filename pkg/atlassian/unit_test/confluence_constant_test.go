package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"testing"
)

func TestClient(t *testing.T) {
	assert.String(t, "no space", constant.NoSpace)
	assert.String(t, "view", constant.ViewFormat)
	assert.String(t, "atlas_doc_format", constant.AtlasFormat)
	assert.String(t, "export_view", constant.ExportFormat)
	assert.String(t, "anonymous_export_view", constant.AnonymousFormat)
	assert.String(t, "styled_view", constant.StyledFormat)
	assert.String(t, "editor", constant.EditFormat)
}
