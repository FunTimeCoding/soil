package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestReconcileToolAnswersPlan(t *testing.T) {
	result := base.New(t, service_tester.WithNetbox()).Tools().MustCallTool(
		constant.Reconcile,
		map[string]any{constant.FullParameter: true},
	)
	assert.StringContains(t, `"hypervisor": "alfa"`, result)
	assert.StringContains(t, `"written": false`, result)
}

func TestReconcileRouteAnswersPlan(t *testing.T) {
	c := base.New(t, service_tester.WithNetbox()).GeneratedClient()
	r, e := c.ReconcileWithResponse(
		t.Context(),
		&client.ReconcileParams{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, r.StatusCode())
	assert.String(t, "alfa", r.JSON200.Hypervisor)
	assert.False(t, r.JSON200.Written)
	unknown, f := c.ReconcileWithResponse(
		t.Context(),
		&client.ReconcileParams{Instance: new("missing")},
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 404, unknown.StatusCode())
}
