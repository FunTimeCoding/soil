package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestUnreachableHypervisorFailsFastWithoutBeingAskedAgain(t *testing.T) {
	o := service_tester.New(t)
	s := o.Service()
	o.Alfa().Hypervisor.SetFailure(
		connection.New(
			errorsConstant.Unreachable,
			"alfa",
			"",
			errorsConstant.Refused,
		),
	)
	_, e := s.ListNodes("alfa")
	assert.True(t, unavailable.Is(e))
	o.Alfa().Hypervisor.SetFailure(nil)
	_, f := s.ListNodes("alfa")
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "proxmox:alfa unreachable since", f.Error())
}

func TestMissingGuestLeavesGuardClosed(t *testing.T) {
	o := service_tester.New(t)
	s := o.Service()
	_, e := s.GetMachine("alfa", 999, "alfa")
	assert.True(t, not_found.Is(e))
	nodes, f := s.ListNodes("alfa")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, nodes)
}
