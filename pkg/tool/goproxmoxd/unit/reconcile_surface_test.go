package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/model_context_tester"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestTheReconcileToolAnswersAPlan(t *testing.T) {
	o := model_context_tester.New(t)
	result := o.Client.MustCallTool(
		constant.Reconcile,
		map[string]any{constant.FullParameter: true},
	)
	assert.StringContains(t, `"hypervisor": "test"`, result)
	assert.StringContains(t, `"written": false`, result)
}

func TestTheReconcileRouteAnswersAPlan(t *testing.T) {
	s := base.New(t)
	c, e := client.NewClientWithResponses(
		s.Address(webConstant.RootPath),
		client.WithHTTPClient(s.Client()),
		client.WithRequestEditorFn(
			web.BearerEditor(generative.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)
	r, e := c.ReconcileWithResponse(
		t.Context(),
		&client.ReconcileParams{Instance: new("test")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.String(t, "test", r.JSON200.Hypervisor)
	assert.False(t, r.JSON200.Written)
	unknown, f := c.ReconcileWithResponse(
		t.Context(),
		&client.ReconcileParams{Instance: new("missing")},
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 404, unknown.StatusCode())
}
