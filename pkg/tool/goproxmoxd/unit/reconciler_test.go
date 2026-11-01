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

func TestTheReconcilerOnlyPromotes(t *testing.T) {
	c, n := reconcileFloor()
	r := reconciler.New(
		service_tester.IntentService(t, c, n, memory.New()),
		constant.ReconcileInterval,
		logger.New(t.Context()),
		memory.New(),
	)
	defer r.Close()
	r.Reconcile()
	assert.String(t, "active", *n.UpdatedMachines["alfa"].Status)
	assert.Count(t, 1, n.UpdatedMachines)
}

func TestAnUnreachableHypervisorIsOnlyLogged(t *testing.T) {
	c, n := reconcileFloor()
	c.SetFailure(
		&net.OpError{
			Op:  "dial",
			Net: "tcp",
			Err: os.NewSyscallError("connect", syscall.EHOSTUNREACH),
		},
	)
	m := memory.New()
	r := reconciler.New(
		service_tester.IntentService(t, c, n, m),
		constant.ReconcileInterval,
		logger.New(t.Context()),
		m,
	)
	defer r.Close()
	r.Reconcile()
	assert.Count(t, 0, m.Events())
	assert.Count(t, 0, n.UpdatedMachines)
}
