package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"strings"
	"testing"
)

func TestAWriteSavesWhileGoquerydIsAway(t *testing.T) {
	o := service_tester.New(t)
	o.Indexer.Failure = unavailable.Format("goqueryd down")
	m, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	stored, f := o.Service.GetMemory(m.Identifier)
	assert.FatalOnError(t, f)
	assert.String(t, "retry policy", stored.Name)
	assert.Count(t, 0, o.Reporter.Events())
}

func TestTheGuardHoldsCallsOffAfterAnOutage(t *testing.T) {
	o := service_tester.New(t)
	o.Indexer.Failure = unavailable.Format("goqueryd down")
	_, e := o.Service.CreateMemory(scopedOption("first", ""))
	assert.FatalOnError(t, e)
	o.Indexer.Failure = nil
	_, f := o.Service.CreateMemory(scopedOption("second", ""))
	assert.FatalOnError(t, f)
	assert.Count(t, 0, o.Indexer.Pushed)
}

func TestReconcileIndexesWhatAWriteMissed(t *testing.T) {
	o := service_tester.New(t)
	o.Indexer.Failure = validation.New("refused")
	m, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	assert.Count(t, 0, o.Indexer.Pushed)
	o.Indexer.Failure = nil
	assert.Count(t, 0, o.Service.Reconcile(t.Context()))
	assert.Count(t, 1, o.Indexer.Pushed)
	assert.String(
		t,
		fmt.Sprintf("memory/%d", m.Identifier),
		o.Indexer.Pushed[0].Path,
	)
}

func TestReconcileUnindexesAForgottenMemory(t *testing.T) {
	o := service_tester.New(t)
	m, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	o.Indexer.Failure = validation.New("refused")
	assert.FatalOnError(t, o.Service.ForgetMemory(m.Identifier, "test"))
	assert.Count(t, 0, o.Indexer.Deleted)
	o.Indexer.Failure = nil
	assert.Count(t, 0, o.Service.Reconcile(t.Context()))
	assert.Count(t, 1, o.Indexer.Deleted)
	assert.String(
		t,
		fmt.Sprintf("memory/%d", m.Identifier),
		o.Indexer.Deleted[0].Path,
	)
}

func TestReconcileStopsAtAnOutage(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service.CreateMemory(scopedOption("first", ""))
	assert.FatalOnError(t, e)
	_, f := o.Service.CreateMemory(scopedOption("second", ""))
	assert.FatalOnError(t, f)
	o.Indexer.Failure = unavailable.Format("goqueryd down")
	failed := o.Service.Reconcile(t.Context())
	assert.Count(t, 1, failed)
	assert.True(t, unavailable.Is(failed[0]))
}

func TestAWriteThatIndexesReportsNothing(t *testing.T) {
	o := service_tester.New(t)
	m, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, o.Service.ForgetMemory(m.Identifier, "test"))
	assert.Count(t, 0, o.Reporter.Events())
}

func TestAWriteReportsAFailureNobodyElseReports(t *testing.T) {
	o := service_tester.New(t)
	o.Indexer.Failure = fmt.Errorf("goqueryd answered nonsense")
	_, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	assert.Count(t, 1, o.Reporter.Events())
}

func TestProfileLeavesTheTopicOutWhileGoquerydIsAway(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	o.Indexer.Failure = unavailable.Format("goqueryd down")
	result, _, f := o.Service.Profile("retry", "", false)
	assert.FatalOnError(t, f)
	notice := "goqueryd is away - topic-matched memories and completions left out"
	assert.String(t, notice, result.Notice)
	assert.True(t, strings.HasPrefix(result.Text, notice))
	assert.Count(t, 1, result.Index)
	assert.Count(t, 0, result.Relevant)
}

func TestProfileReportsAFailureThatIsNoOutageAndStillAnswers(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	o.Indexer.Failure = fmt.Errorf("goqueryd answered nonsense")
	result, _, f := o.Service.Profile("retry", "", false)
	assert.FatalOnError(t, f)
	assert.String(
		t,
		"goqueryd failed - topic-matched memories and completions left out",
		result.Notice,
	)
	assert.Count(t, 1, result.Index)
	assert.Count(t, 2, o.Reporter.Events())
}
