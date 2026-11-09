package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"testing"
)

func TestMockAnswersUnknownMachineAsNotFound(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	_, e := m.GetMachine(t.Context(), 999, &client.GetMachineParams{})
	assert.True(t, not_found.Is(e))
}

func TestMockRefusesWriteToMissingMachine(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	_, e := m.StartMachine(t.Context(), 999, &client.StartMachineParams{})
	assert.True(t, not_found.Is(e))
	_, f := m.CreateMachineSnapshot(
		t.Context(),
		999,
		&client.CreateMachineSnapshotParams{},
		client.CreateMachineSnapshotJSONRequestBody{Name: "alfa"},
	)
	assert.True(t, not_found.Is(f))
}

func TestMockStartsMachineItCreated(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	created, e := m.CreateMachine(
		t.Context(),
		&client.CreateMachineParams{},
		client.CreateMachineJSONRequestBody{Name: "bravo", Node: "charlie"},
	)
	assert.FatalOnError(t, e)
	_, f := m.StartMachine(
		t.Context(),
		created.Identifier,
		&client.StartMachineParams{},
	)
	assert.FatalOnError(t, f)
	v, g := m.GetMachine(
		t.Context(),
		created.Identifier,
		&client.GetMachineParams{},
	)
	assert.FatalOnError(t, g)
	assert.String(t, "running", *v.Status)
}
