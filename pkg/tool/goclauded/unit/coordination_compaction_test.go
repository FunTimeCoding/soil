package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestCompactionThroughClientReachesWatcher(t *testing.T) {
	s := base.New(t)
	watcher := s.NewSession(t)
	watcher.Announce(watcher.Name(), "leading")
	watcher.CheckLive()
	_, e := watcher.RestClient.PostSessionLabelWithResponse(
		watcher.Context,
		watcher.UUID,
		labelBody(
			constant.CompactionWatchLabel,
			constant.CompactionWatchValue,
			watcher.Name(),
		),
	)
	assert.FatalOnError(t, e)
	compacting := s.NewSession(t)
	compacting.Announce(compacting.Name(), "building")
	compacting.CheckLive()
	assert.FatalOnError(
		t,
		s.Client(t).PostCompaction(
			compacting.Context,
			client.PostCompactionJSONRequestBody{
				Session: compacting.UUID,
				Phase:   client.Started,
			},
		),
	)
	notifications := clientKindEntries(
		watcher.CheckLive().Entries,
		constant.QueueNotification,
	)
	assert.Count(t, 1, notifications)
	assert.String(
		t,
		join.Empty(
			"compaction: ",
			compacting.Name(),
			" started compacting - building",
		),
		notifications[0].Body,
	)
}

func TestCompactionThroughClientRefusesUnknownPhase(t *testing.T) {
	s := base.New(t)
	compacting := s.NewSession(t)
	compacting.CheckLive()
	e := s.Client(t).PostCompaction(
		compacting.Context,
		client.PostCompactionJSONRequestBody{
			Session: compacting.UUID,
			Phase:   "halfway",
		},
	)
	assert.True(t, validation.Is(e))
}
