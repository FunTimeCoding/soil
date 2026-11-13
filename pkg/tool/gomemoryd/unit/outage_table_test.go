package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOperationsOnQueryOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outageServer(t, r)
		},
		spec,
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			"GetProfile": webConstant.OutageDegrades,
			"PostMemory": webConstant.OutageDegrades,
		},
		nil,
		nil,
	)
}

func TestToolsOnQueryOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outageServer(t, r)
		},
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			"profile":     webConstant.OutageDegrades,
			"save_memory": webConstant.OutageDegrades,
		},
		nil,
	)
}
