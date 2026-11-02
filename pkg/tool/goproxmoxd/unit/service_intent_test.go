package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestAStartedMachineIsRecordedActive(t *testing.T) {
	c := mock_client.New()
	c.AddNode("test")
	c.AddMachine("test", 101, "alfa")
	v, n := service_tester.IntentService(t, c, memory.New())
	_, e := v.StartMachine("test", 101, "")
	assert.FatalOnError(t, e)
	assert.String(t, "active", *n.UpdatedMachines["alfa"].Status)
}

func TestAShutDownContainerIsRecordedOffline(t *testing.T) {
	c := mock_client.New()
	c.AddNode("test")
	c.AddContainer("test", 201, "bravo")
	v, n := service_tester.IntentService(t, c, memory.New())
	_, e := v.ShutdownContainer("test", 201, "")
	assert.FatalOnError(t, e)
	assert.String(t, "offline", *n.UpdatedMachines["bravo"].Status)
}

func TestAnUpwardReconcileOnlyPromotes(t *testing.T) {
	c, machines := reconcileFloor()
	v, n := service_tester.IntentService(t, c, memory.New(), machines...)
	p, e := v.Reconcile(context.Background(), "test", false, false)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, p.Changes)
	assert.String(t, "alfa", p.Changes[0].Name)
	assert.String(t, "offline", p.Changes[0].From)
	assert.String(t, "active", p.Changes[0].To)
	assert.Strings(t, []string{"delta"}, p.Unknown)
	assert.False(t, p.Written)
	assert.Count(t, 0, n.UpdatedMachines)
}

func TestAFullReconcileAlsoDemotesAndWrites(t *testing.T) {
	c, machines := reconcileFloor()
	v, n := service_tester.IntentService(t, c, memory.New(), machines...)
	p, e := v.Reconcile(context.Background(), "test", true, true)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, p.Changes)
	assert.String(t, "alfa", p.Changes[0].Name)
	assert.String(t, "active", p.Changes[0].To)
	assert.String(t, "bravo", p.Changes[1].Name)
	assert.String(t, "offline", p.Changes[1].To)
	assert.True(t, p.Written)
	assert.String(t, "active", *n.UpdatedMachines["alfa"].Status)
	assert.String(t, "offline", *n.UpdatedMachines["bravo"].Status)
	assert.Count(t, 2, n.UpdatedMachines)
}

func TestAReconcileWithoutTheSwitchIsNotConfigured(t *testing.T) {
	c, _ := reconcileFloor()
	_, e := service_tester.ContainerService(c).Reconcile(
		context.Background(),
		"test",
		true,
		false,
	)
	assert.True(t, not_configured.Is(e))
}
