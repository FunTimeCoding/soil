package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"testing"
)

func TestSnippetRoutesAnswerMissingSnippetNotFound(t *testing.T) {
	c := base.New(t).GeneratedClient()
	got, e := c.GetSnippetWithResponse(
		t.Context(),
		"alfa.yaml",
		&client.GetSnippetParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 404, got.StatusCode())
	removed, e := c.DeleteSnippetWithResponse(
		t.Context(),
		"alfa.yaml",
		&client.DeleteSnippetParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 404, removed.StatusCode())
}
