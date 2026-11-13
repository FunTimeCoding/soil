package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOperationsOnEmbedderOutage(t *testing.T) {
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
			"GetSearch":    webConstant.OutageReaches,
			"PostDocument": webConstant.OutageDegrades,
		},
		nil,
		nil,
	)
}

func TestToolsOnEmbedderOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outageServer(t, r)
		},
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			"push":   webConstant.OutageDegrades,
			"search": webConstant.OutageReaches,
		},
		nil,
	)
}
