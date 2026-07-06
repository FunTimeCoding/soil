package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestClient(t *testing.T) {
	assert.String(t, "no space", NoSpace)
	assert.String(t, "view", ViewFormat)
	assert.String(t, "atlas_doc_format", AtlasFormat)
	assert.String(t, "export_view", ExportFormat)
	assert.String(t, "anonymous_export_view", AnonymousFormat)
	assert.String(t, "styled_view", StyledFormat)
	assert.String(t, "editor", EditFormat)
}
