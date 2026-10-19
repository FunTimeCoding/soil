package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	toolConstant "github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/model_context_tester"
	"testing"
)

func TestAStoreFaultCarriesItsOwnTextOverModelContext(t *testing.T) {
	o := model_context_tester.New(t)
	o.Store().Close()
	assert.StringContains(
		t,
		`"error": "sql: database is closed"`,
		o.Client.MustCallToolError(
			toolConstant.GetAlerts,
			map[string]any{constant.ParameterName: "HighMemory"},
		),
	)
}
