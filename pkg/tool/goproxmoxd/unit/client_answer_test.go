package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/client_tester"
	"testing"
)

func TestClientListsHypervisorsMachines(t *testing.T) {
	c, s := client_tester.New(t)
	s.MockClient.AddMachine("test", 101, "alfa")
	v, e := c.ListMachines(
		t.Context(),
		&client.ListMachinesParams{Instance: new("test")},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "alfa", v[0].Name)
}

func TestClientAnswersUnknownMachineAsNotFound(t *testing.T) {
	c, _ := client_tester.New(t)
	_, e := c.GetMachine(
		t.Context(),
		999,
		&client.GetMachineParams{Instance: new("test")},
	)
	assert.True(t, not_found.Is(e))
}

func TestClientAnswersUnreachableHypervisorAsUnavailable(t *testing.T) {
	c, s := client_tester.New(t)
	s.MockClient.SetFailure(
		connection.New(
			errorsConstant.Unreachable,
			"test",
			"",
			errorsConstant.Refused,
		),
	)
	_, e := c.ListNodes(
		t.Context(),
		&client.ListNodesParams{Instance: new("test")},
	)
	assert.True(t, unavailable.Is(e))
}
