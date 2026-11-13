package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOperationsOnForgeOutage(t *testing.T) {
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
			"PostPublication": webConstant.OutageReaches,
		},
		nil,
		map[string]string{
			"PostAuthority": string(notation.Marshal(server_tester.RootBody())),
		},
	)
}

func TestToolsOnForgeOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outageServer(t, r)
		},
		webConstant.OutageElsewhere,
		map[string]webConstant.OutageReach{"publish": webConstant.OutageReaches},
		nil,
	)
}
