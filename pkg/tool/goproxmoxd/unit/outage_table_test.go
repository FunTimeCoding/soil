package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/outage_server"
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
			"CreateSnippet": webConstant.OutageElsewhere,
			"DeleteSnippet": webConstant.OutageElsewhere,
			"GetSnippet":    webConstant.OutageElsewhere,
			"ListSnippets":  webConstant.OutageElsewhere,
		},
		nil,
		map[string]string{"UpdateMachine": `{"name":"alfa"}`},
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
			"Reconcile": webConstant.OutageReaches,
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
		map[string]webConstant.OutageReach{
			constant.CreateSnippet: webConstant.OutageElsewhere,
			constant.DeleteSnippet: webConstant.OutageElsewhere,
			constant.GetSnippet:    webConstant.OutageElsewhere,
			constant.ListSnippets:  webConstant.OutageElsewhere,
		},
		map[string]map[string]any{constant.UpdateMachine: {"name": "alfa"}},
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
			constant.Reconcile: webConstant.OutageReaches,
		},
		nil,
	)
}
