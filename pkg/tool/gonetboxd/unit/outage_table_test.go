package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOperationsOnNetBoxOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	address := `{"address":"00:00:5e:00:53:01"}`
	outage_tester.Operations(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outageServer(t, r)
		},
		spec,
		webConstant.OutageReaches,
		nil,
		nil,
		map[string]string{
			"SetInterfacePhysicalAddress":        address,
			"SetVirtualInterfacePhysicalAddress": address,
		},
	)
}

func TestToolsOnNetBoxOutage(t *testing.T) {
	outage_tester.Tools(
		t,
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outageServer(t, r)
		},
		webConstant.OutageReaches,
		nil,
		map[string]map[string]any{"update_journal_entry": {"comments": "alfa"}},
	)
}
