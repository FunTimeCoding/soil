package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"testing"
)

func TestClientListsHypervisorsMachines(t *testing.T) {
	s := base.New(t)
	c := s.ProxmoxClient()
	s.Alfa().Hypervisor.AddMachine("alfa", 101, "alfa")
	v, e := c.ListMachines(
		t.Context(),
		&client.ListMachinesParams{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "alfa", v[0].Name)
}

func TestClientAnswersUnknownMachineAsNotFound(t *testing.T) {
	c := base.New(t).ProxmoxClient()
	_, e := c.GetMachine(
		t.Context(),
		999,
		&client.GetMachineParams{Instance: new("alfa")},
	)
	assert.True(t, not_found.Is(e))
}

func TestClientAnswersUnreachableHypervisorAsUnavailable(t *testing.T) {
	s := base.New(t)
	c := s.ProxmoxClient()
	s.Alfa().Hypervisor.SetFailure(
		connection.New(
			errorsConstant.Unreachable,
			"alfa",
			"",
			errorsConstant.Refused,
		),
	)
	_, e := c.ListNodes(
		t.Context(),
		&client.ListNodesParams{Instance: new("alfa")},
	)
	assert.True(t, unavailable.Is(e))
}
