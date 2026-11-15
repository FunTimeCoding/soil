package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/unit/outage_server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestToolsOnInstanceOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r)
		},
		webConstant.OutageReaches,
		nil,
		map[string]map[string]any{constant.QueryRange: {"start": "1h"}},
	)
}
