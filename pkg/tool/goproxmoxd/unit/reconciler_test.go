package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/reconciler"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestReconcilerOnlyPromotes(t *testing.T) {
	o := service_tester.New(t, service_tester.WithNetbox())
	o.ReconcileFloor()
	r := reconciler.New(
		o.Service(),
		o.Inventory(),
		constant.ReconcileInterval,
		logger.New(t.Context()),
		memory.New(),
	)
	defer r.Close()
	r.Reconcile()
	assert.String(t, "active", *o.Netbox.UpdatedMachines["alfa"].Status)
	assert.Count(t, 1, o.Netbox.UpdatedMachines)
}

func TestUnreachableHypervisorIsOnlyLogged(t *testing.T) {
	m := memory.New()
	o := service_tester.New(
		t,
		service_tester.WithNetbox(),
		service_tester.WithReporter(m),
	)
	o.ReconcileFloor()
	o.Alfa().Hypervisor.SetFailure(
		&net.OpError{
			Op:  "dial",
			Net: "tcp",
			Err: os.NewSyscallError("connect", syscall.EHOSTUNREACH),
		},
	)
	r := reconciler.New(
		o.Service(),
		o.Inventory(),
		constant.ReconcileInterval,
		logger.New(t.Context()),
		m,
	)
	defer r.Close()
	r.Reconcile()
	assert.Count(t, 0, m.Events())
	assert.Count(t, 0, o.Netbox.UpdatedMachines)
}
