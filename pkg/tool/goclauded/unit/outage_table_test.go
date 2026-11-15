package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/outage_server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOperationsOnMemoryOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.MemoryPeer)
		},
		spec,
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			"GetTimeline": webConstant.OutageDegrades,
		},
		nil,
		nil,
	)
}

func TestOperationsOnQueryOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.QueryPeer)
		},
		spec,
		webConstant.OutageElsewhere,
		nil,
		nil,
		nil,
	)
}

func TestOperationsOnUtilizationOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.UtilizationPeer)
		},
		spec,
		webConstant.OutageElsewhere,
		nil,
		nil,
		nil,
	)
}

func TestToolsOnMemoryOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.MemoryPeer)
		},
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			constant.History: webConstant.OutageDegrades,
		},
		nil,
	)
}

func TestToolsOnQueryOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.QueryPeer)
		},
		webConstant.OutageElsewhere,
		nil,
		nil,
	)
}

func TestToolsOnUtilizationOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.UtilizationPeer)
		},
		webConstant.OutageElsewhere,
		nil,
		nil,
	)
}
