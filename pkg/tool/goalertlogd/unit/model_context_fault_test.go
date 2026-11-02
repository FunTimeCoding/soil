package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/model_context_tester"
	"testing"
)

func TestAStoreFaultCarriesItsOwnTextOverModelContext(t *testing.T) {
	o := model_context_tester.New(t)
	o.Daemon().Close()
	assert.StringContains(
		t,
		`"error": "sql: database is closed"`,
		o.Client.MustCallToolError(
			constant.GetAlerts,
			map[string]any{generativeConstant.ParameterName: "HighMemory"},
		),
	)
}
