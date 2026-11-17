package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestSameNamedGuestIsNeverReconciled(t *testing.T) {
	o := service_tester.New(t, service_tester.WithNetbox())
	o.AmbiguousFloor()
	v := o.Service()
	p, e := v.Reconcile(context.Background(), "alfa", true, true)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"alfa"}, p.Ambiguous)
	assert.Count(t, 0, p.Changes)
	assert.Count(t, 0, o.Netbox.UpdatedMachines)
}

func TestSameNamedGuestIsNeverRecordedOnStart(t *testing.T) {
	m := memory.New()
	o := service_tester.New(
		t,
		service_tester.WithNetbox(),
		service_tester.WithReporter(m),
	)
	o.AmbiguousFloor()
	v := o.Service()
	_, e := v.StartMachine("alfa", 101, "")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, o.Netbox.UpdatedMachines)
	assert.Count(t, 0, m.Events())
}
