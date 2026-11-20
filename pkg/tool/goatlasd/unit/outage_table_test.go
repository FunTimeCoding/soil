package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/outage_server"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	gooutpostdConstant "github.com/funtimecoding/soil/pkg/tool/gooutpostd/constant"
	gopnsensedConstant "github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	goprocessdConstant "github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOperationsOnNetboxOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.New(t).Operations(
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, gonetboxdConstant.Identity.Name())
		},
		spec,
		webConstant.OutageElsewhere,
		nil,
		nil,
		nil,
	)
}

func TestOperationsOnKubernetesOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.New(t).Operations(
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, constant.SourceKubernetes)
		},
		spec,
		webConstant.OutageElsewhere,
		nil,
		nil,
		nil,
	)
}

func TestOperationsOnProcessOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.New(t).Operations(
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, goprocessdConstant.Identity.Name())
		},
		spec,
		webConstant.OutageElsewhere,
		nil,
		nil,
		nil,
	)
}

func TestOperationsOnOutpostOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.New(t).Operations(
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(
				t,
				r,
				join.Colon(
					gooutpostdConstant.Identity.Name(),
					constant.FixtureVirtualMachineNode,
				),
			)
		},
		spec,
		webConstant.OutageElsewhere,
		nil,
		nil,
		nil,
	)
}

func TestOperationsOnLeaseOutage(t *testing.T) {
	spec, e := server.GetSpec()
	assert.FatalOnError(t, e)
	outage_tester.New(t).Operations(
		func(r *memory.Memory) (*model_context_server.Server, func() error) {
			return outage_server.New(t, r, gopnsensedConstant.Identity.Name())
		},
		spec,
		webConstant.OutageElsewhere,
		nil,
		nil,
		nil,
	)
}
