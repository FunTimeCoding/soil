package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/outage_server"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOperationsOnInstanceOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(
				t,
				r,
				join.Colon(constant.InstancePeer, "alfa"),
			)
		},
		spec,
		webConstant.OutageReaches,
		map[string]webConstant.OutageReach{
			"ListInstances": webConstant.OutageElsewhere,
		},
		nil,
		nil,
	)
}

func TestOperationsOnNetboxOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, gonetboxdConstant.Identity.Name())
		},
		spec,
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			"RequestPower": webConstant.OutageDegrades,
		},
		nil,
		nil,
	)
}

func TestToolsOnInstanceOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(
				t,
				r,
				join.Colon(constant.InstancePeer, "alfa"),
			)
		},
		webConstant.OutageReaches,
		nil,
		nil,
	)
}

func TestToolsOnNetboxOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, gonetboxdConstant.Identity.Name())
		},
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{
			constant.Start:    webConstant.OutageDegrades,
			constant.Shutdown: webConstant.OutageDegrades,
		},
		nil,
	)
}
