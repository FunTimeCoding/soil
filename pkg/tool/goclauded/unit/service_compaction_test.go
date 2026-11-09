package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestFinishedCompactionWakesWatcher(t *testing.T) {
	s := service_tester.New(t)
	s.Watch("watcher")
	compacting := s.Check("compacting").Callsign
	s.Announce("compacting", compacting, "renaming the store", "")
	s.Compact("compacting", constant.CompactionFinished, "auto")
	drained, e := s.Message.DrainImmediateQueue("watcher", "")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, drained)
	assert.String(
		t,
		join.Empty(
			"compaction: ",
			compacting,
			" finished compacting (auto) - renaming the store",
		),
		drained[0].Body,
	)
}

func TestLeadAlsoOnCallIsTold(t *testing.T) {
	s := service_tester.New(t)
	s.Check("watcher")
	s.SetLabel("watcher", constant.CompactionWatchLabel, "on_call,lead")
	s.Check("compacting")
	s.Compact("compacting", constant.CompactionFinished, "auto")
	assert.Integer(t, 1, s.Channel("watcher"))
}

func TestOnCallAloneIsNotLead(t *testing.T) {
	s := service_tester.New(t)
	s.Check("watcher")
	s.SetLabel("watcher", constant.CompactionWatchLabel, "on_call")
	s.Check("compacting")
	s.Compact("compacting", constant.CompactionFinished, "auto")
	assert.Integer(t, 0, s.Channel("watcher"))
}

func TestStartedCompactionWaitsForWatchersPrompt(t *testing.T) {
	s := service_tester.New(t)
	s.Watch("watcher")
	s.Check("compacting")
	s.Compact("compacting", constant.CompactionStarted, "manual")
	assert.Integer(t, 0, s.Channel("watcher"))
	assert.Count(
		t,
		1,
		fixture.KindEntries(
			s.Check("watcher").Entries,
			constant.QueueNotification,
		),
	)
}

func TestCompactionReachesEveryWatcherAndNobodyElse(t *testing.T) {
	s := service_tester.New(t)
	s.Watch("first")
	s.Watch("second")
	s.Check("bystander")
	s.Check("compacting")
	s.Compact("compacting", constant.CompactionFinished, "auto")
	assert.Integer(t, 1, s.Channel("first"))
	assert.Integer(t, 1, s.Channel("second"))
	assert.Integer(t, 0, s.Channel("bystander"))
	assert.Integer(t, 0, s.Channel("compacting"))
}

func TestWatcherIsNotToldOfItsOwnCompaction(t *testing.T) {
	s := service_tester.New(t)
	s.Watch("watcher")
	s.Compact("watcher", constant.CompactionFinished, "auto")
	assert.Integer(t, 0, s.Channel("watcher"))
	assert.Count(
		t,
		0,
		fixture.KindEntries(
			s.Check("watcher").Entries,
			constant.QueueNotification,
		),
	)
}

func TestCompactionWithoutWatchersIsStillLogged(t *testing.T) {
	s := service_tester.New(t)
	compacting := s.Check("compacting").Callsign
	s.Compact("compacting", constant.CompactionStarted, "auto")
	events := s.Store.Events(
		event_query.New().Kind(constant.Compaction).SetLimit(10),
	)
	assert.Count(t, 1, events)
	assert.String(t, compacting, events[0].Actor)
	assert.String(t, "started", events[0].Metadata[constant.Phase])
	assert.String(t, "auto", events[0].Metadata[constant.Trigger])
}

func TestUnknownCompactionPhaseIsRefused(t *testing.T) {
	s := service_tester.New(t)
	s.Watch("watcher")
	s.Check("compacting")
	e := s.Roster.RecordCompaction("compacting", "halfway", "auto")
	assert.True(t, validation.Is(e))
	assert.Integer(t, 0, s.Channel("watcher"))
}

func TestCompactionOfUnknownSessionIsIgnored(t *testing.T) {
	s := service_tester.New(t)
	s.Watch("watcher")
	s.Compact("stranger", constant.CompactionFinished, "auto")
	assert.Integer(t, 0, s.Channel("watcher"))
}
