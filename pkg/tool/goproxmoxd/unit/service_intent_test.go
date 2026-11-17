package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestStartedMachineIsRecordedActive(t *testing.T) {
	o := service_tester.New(t, service_tester.WithNetbox())
	o.Alfa().Hypervisor.AddMachine("alfa", 101, "alfa")
	v := o.Service()
	_, e := v.StartMachine("alfa", 101, "")
	assert.FatalOnError(t, e)
	assert.String(t, "active", *o.Netbox.UpdatedMachines["alfa"].Status)
}

func TestShutDownContainerIsRecordedOffline(t *testing.T) {
	o := service_tester.New(t, service_tester.WithNetbox())
	o.Alfa().Hypervisor.AddContainer("alfa", 201, "bravo")
	v := o.Service()
	_, e := v.ShutdownContainer("alfa", 201, "")
	assert.FatalOnError(t, e)
	assert.String(t, "offline", *o.Netbox.UpdatedMachines["bravo"].Status)
}

func TestUpwardReconcileOnlyPromotes(t *testing.T) {
	o := service_tester.New(t, service_tester.WithNetbox())
	o.ReconcileFloor()
	v := o.Service()
	p, e := v.Reconcile(context.Background(), "alfa", false, false)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, p.Changes)
	assert.String(t, "alfa", p.Changes[0].Name)
	assert.String(t, "offline", p.Changes[0].From)
	assert.String(t, "active", p.Changes[0].To)
	assert.Strings(t, []string{"delta"}, p.Unknown)
	assert.False(t, p.Written)
	assert.Count(t, 0, o.Netbox.UpdatedMachines)
}

func TestFullReconcileAlsoDemotesAndWrites(t *testing.T) {
	o := service_tester.New(t, service_tester.WithNetbox())
	o.ReconcileFloor()
	v := o.Service()
	p, e := v.Reconcile(context.Background(), "alfa", true, true)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, p.Changes)
	assert.String(t, "alfa", p.Changes[0].Name)
	assert.String(t, "active", p.Changes[0].To)
	assert.String(t, "bravo", p.Changes[1].Name)
	assert.String(t, "offline", p.Changes[1].To)
	assert.True(t, p.Written)
	assert.String(t, "active", *o.Netbox.UpdatedMachines["alfa"].Status)
	assert.String(t, "offline", *o.Netbox.UpdatedMachines["bravo"].Status)
	assert.Count(t, 2, o.Netbox.UpdatedMachines)
}

func TestReconcileWithoutSwitchIsNotConfigured(t *testing.T) {
	o := service_tester.New(t)
	o.ReconcileFloor()
	_, e := o.Service().Reconcile(context.Background(), "alfa", true, false)
	assert.True(t, not_configured.Is(e))
}
