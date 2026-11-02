package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestTheSnippetRoutesAnswerAMissingSnippetNotFound(t *testing.T) {
	s := base.New(t)
	c, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)
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
