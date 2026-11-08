package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestTheCensusListsEveryMethodOfTheType(t *testing.T) {
	d := prepareCensus(t)
	c, e := service_tester.NewService(t).Census.Census(
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
	q := censusMethod(t, c, "Quiet")
	assert.Count(t, 0, q.Callers)
	assert.Count(t, 0, q.Uses)
	assert.Count(t, 0, q.Reach.Written)
}

func TestTheCensusCountsCallersByKind(t *testing.T) {
	d := prepareCensus(t)
	c, e := service_tester.NewService(t).Census.Census(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	place := censusMethod(t, c, "Place").Callers
	assert.Count(t, 3, place)
	assert.String(
		t,
		"production pkg/model_context Place 1",
		callerLine(place[0]),
	)
	assert.String(t, "production pkg/worker Tick 1", callerLine(place[1]))
	assert.String(t, "test pkg/unit TestPlace 1", callerLine(place[2]))
	now := censusMethod(t, c, "Now").Callers
	assert.Count(t, 1, now)
	assert.String(
		t,
		"support pkg/unit/service_tester Build 1",
		callerLine(now[0]),
	)
	step := censusMethod(t, c, "step").Callers
	assert.Count(t, 1, step)
	assert.String(
		t,
		"internal pkg/service Service.Settle 1",
		callerLine(step[0]),
	)
}

func TestTheCensusReadsTheReceiverThroughFreeFunctions(t *testing.T) {
	d := prepareCensus(t)
	c, e := service_tester.NewService(t).Census.Census(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	place := censusMethod(t, c, "Place")
	assert.Strings(t, []string{"store"}, place.Fields)
	assert.Strings(t, []string{"record"}, place.Calls)
	assert.Count(t, 1, place.Uses)
	assert.String(t, "store.Store", place.Uses[0].Type)
	assert.Strings(t, []string{"Order"}, place.Uses[0].Methods)
	now := censusMethod(t, c, "Now")
	assert.Strings(t, []string{"stamp"}, now.Functions)
	assert.Strings(t, []string{"clock"}, now.Fields)
}

func TestTheCensusDerivesTablesFromTheGormCalls(t *testing.T) {
	d := prepareCensus(t)
	c, e := service_tester.NewService(t).Census.Census(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	place := censusMethod(t, c, "Place")
	assert.Strings(t, []string{"orders"}, place.Tables.Read)
	assert.Count(t, 0, place.Tables.Written)
	assert.Strings(t, []string{"orders"}, place.Reach.Read)
	assert.Strings(t, []string{"sent_receipts"}, place.Reach.Written)
	tally := censusMethod(t, c, "Tally")
	assert.Strings(t, []string{"orders"}, tally.Tables.Read)
	record := censusMethod(t, c, "record")
	assert.Strings(t, []string{"sent_receipts"}, record.Tables.Written)
	assert.Strings(
		t,
		[]string{"pkg/model/receipt", "pkg/store"},
		record.Packages,
	)
}

func TestTheCensusMarksTheTransaction(t *testing.T) {
	d := prepareCensus(t)
	c, e := service_tester.NewService(t).Census.Census(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	settle := censusMethod(t, c, "Settle")
	assert.True(t, settle.Atomically)
	assert.False(t, settle.TransactionStep)
	assert.Strings(t, []string{"step"}, settle.Calls)
	assert.Strings(t, []string{"sent_receipts"}, settle.Reach.Written)
	step := censusMethod(t, c, "step")
	assert.True(t, step.TransactionStep)
	assert.False(t, step.Atomically)
	assert.Strings(t, []string{"sent_receipts"}, step.Tables.Written)
}

func TestTheCensusRefusesATypeThatIsNotAStruct(t *testing.T) {
	d := prepareCensus(t)
	_, e := service_tester.NewService(t).Census.Census(
		d,
		"example/pkg/service",
		"Nothing",
	)
	assert.True(t, validation.Is(e))
	assert.String(t, "no struct type Nothing in example/pkg/service", e.Error())
}
