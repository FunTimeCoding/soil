package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"testing"
)

func TestAnUnknownMemoryIsAnsweredNotFound(t *testing.T) {
	s := base.New(t)
	c, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)
	r, f := c.DeleteMemoryWithResponse(t.Context(), 999, nil)
	assert.FatalOnError(t, f)
	assert.Integer(t, http.StatusNotFound, r.StatusCode())
	assert.String(t, "memory not found: 999", r.JSON404.Error)
	assert.Count(t, 0, s.Reporter.Events())
}
