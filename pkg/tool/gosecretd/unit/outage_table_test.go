package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOperationsOnOpenbaoOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outageServer(t, r)
		},
		spec,
		webConstant.OutageReaches,
		nil,
		nil,
		map[string]string{"PutSecret": `{"values":{"alfa":"bravo"}}`},
	)
}

func TestToolsOnOpenbaoOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outageServer(t, r)
		},
		webConstant.OutageReaches,
		nil,
		map[string]map[string]any{
			"put_secret": {"values": map[string]any{"alfa": "bravo"}},
		},
	)
}
