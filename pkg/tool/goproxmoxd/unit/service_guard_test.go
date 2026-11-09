package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/proxmox/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestUnreachableHypervisorFailsFastWithoutBeingAskedAgain(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.AddNode("test")
	s := service_tester.ContainerService(t, c)
	c.SetFailure(
		connection.New(
			errorsConstant.Unreachable,
			"test",
			"",
			errorsConstant.Refused,
		),
	)
	_, e := s.ListNodes("test")
	assert.True(t, unavailable.Is(e))
	c.SetFailure(nil)
	_, f := s.ListNodes("test")
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "proxmox:test unreachable since", f.Error())
}

func TestMissingGuestLeavesGuardClosed(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.AddNode("test")
	s := service_tester.ContainerService(t, c)
	_, e := s.GetMachine("test", 999, "test")
	assert.True(t, not_found.Is(e))
	nodes, f := s.ListNodes("test")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, nodes)
}
