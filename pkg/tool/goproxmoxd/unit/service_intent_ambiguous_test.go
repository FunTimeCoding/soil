package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestASameNamedGuestIsNeverReconciled(t *testing.T) {
	c, n := ambiguousFloor()
	p, e := service_tester.IntentService(t, c, n, memory.New()).Reconcile(
		context.Background(),
		"test",
		true,
		true,
	)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"alfa"}, p.Ambiguous)
	assert.Count(t, 0, p.Changes)
	assert.Count(t, 0, n.UpdatedMachines)
}

func TestASameNamedGuestIsNeverRecordedOnStart(t *testing.T) {
	c, n := ambiguousFloor()
	m := memory.New()
	_, e := service_tester.IntentService(t, c, n, m).StartMachine(
		"test",
		101,
		"",
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, n.UpdatedMachines)
	assert.Count(t, 0, m.Events())
}
