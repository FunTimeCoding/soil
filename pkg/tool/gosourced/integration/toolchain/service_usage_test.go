package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/usage_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestUsageListsEveryMethodOfType(t *testing.T) {
	d := usage_tester.Prepare(t)
	c, e := service_tester.NewService(t).Usage.Usage(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"store", "clock"}, c.Fields)
	var names []string

	for _, m := range c.Methods {
		names = append(names, m.Name)
	}

	assert.Strings(
		t,
		[]string{"Now", "Place", "Quiet", "record", "Settle", "step", "Tally"},
		names,
	)
	q := usage_tester.Method(t, c, "Quiet")
	assert.Count(t, 0, q.Callers)
	assert.Count(t, 0, q.Uses)
	assert.Count(t, 0, q.Reach.Written)
}

func TestUsageCountsCallersByKind(t *testing.T) {
	d := usage_tester.Prepare(t)
	c, e := service_tester.NewService(t).Usage.Usage(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	place := usage_tester.Method(t, c, "Place").Callers
	assert.Count(t, 3, place)
	assert.String(
		t,
		"production pkg/model_context Place 1",
		usage_tester.CallerLine(place[0]),
	)
	assert.String(
		t,
		"production pkg/worker Tick 1",
		usage_tester.CallerLine(place[1]),
	)
	assert.String(
		t,
		"test pkg/unit TestPlace 1",
		usage_tester.CallerLine(place[2]),
	)
	now := usage_tester.Method(t, c, "Now").Callers
	assert.Count(t, 1, now)
	assert.String(
		t,
		"support pkg/unit/service_tester Build 1",
		usage_tester.CallerLine(now[0]),
	)
	step := usage_tester.Method(t, c, "step").Callers
	assert.Count(t, 1, step)
	assert.String(
		t,
		"internal pkg/service Service.Settle 1",
		usage_tester.CallerLine(step[0]),
	)
}

func TestUsageReadsReceiverThroughFreeFunctions(t *testing.T) {
	d := usage_tester.Prepare(t)
	c, e := service_tester.NewService(t).Usage.Usage(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	place := usage_tester.Method(t, c, "Place")
	assert.Strings(t, []string{"store"}, place.Fields)
	assert.Strings(t, []string{"record"}, place.Calls)
	assert.Count(t, 1, place.Uses)
	assert.String(t, "store.Store", place.Uses[0].Type)
	assert.Strings(t, []string{"Order"}, place.Uses[0].Methods)
	now := usage_tester.Method(t, c, "Now")
	assert.Strings(t, []string{"stamp"}, now.Functions)
	assert.Strings(t, []string{"clock"}, now.Fields)
}

func TestUsageDerivesTablesFromGormCalls(t *testing.T) {
	d := usage_tester.Prepare(t)
	c, e := service_tester.NewService(t).Usage.Usage(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	place := usage_tester.Method(t, c, "Place")
	assert.Strings(t, []string{"orders"}, place.Tables.Read)
	assert.Count(t, 0, place.Tables.Written)
	assert.Strings(t, []string{"orders"}, place.Reach.Read)
	assert.Strings(t, []string{"sent_receipts"}, place.Reach.Written)
	tally := usage_tester.Method(t, c, "Tally")
	assert.Strings(t, []string{"orders"}, tally.Tables.Read)
	record := usage_tester.Method(t, c, "record")
	assert.Strings(t, []string{"sent_receipts"}, record.Tables.Written)
	assert.Strings(
		t,
		[]string{"pkg/model/receipt", "pkg/store"},
		record.Packages,
	)
}

func TestUsageMarksTransaction(t *testing.T) {
	d := usage_tester.Prepare(t)
	c, e := service_tester.NewService(t).Usage.Usage(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	settle := usage_tester.Method(t, c, "Settle")
	assert.True(t, settle.Atomically)
	assert.False(t, settle.TransactionStep)
	assert.Strings(t, []string{"step"}, settle.Calls)
	assert.Strings(t, []string{"sent_receipts"}, settle.Reach.Written)
	step := usage_tester.Method(t, c, "step")
	assert.True(t, step.TransactionStep)
	assert.False(t, step.Atomically)
	assert.Strings(t, []string{"sent_receipts"}, step.Tables.Written)
}

func TestUsageFailsOnTypeOtherThanStruct(t *testing.T) {
	d := usage_tester.Prepare(t)
	_, e := service_tester.NewService(t).Usage.Usage(
		d,
		"example/pkg/service",
		"Nothing",
	)
	assert.True(t, validation.Is(e))
	assert.String(t, "no struct type Nothing in example/pkg/service", e.Error())
}
