package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/delivery_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/label_tester"
	"testing"
)

func TestCompactionThroughClientReachesWatcher(t *testing.T) {
	s := base.New(t)
	watcher := s.NewSession()
	watcher.Announce(watcher.Name(), "leading")
	watcher.CheckLive()
	_, e := watcher.GeneratedClient.PostSessionLabelWithResponse(
		watcher.Context,
		watcher.UUID,
		label_tester.Body(
			constant.CompactionWatchLabel,
			constant.CompactionWatchValue,
			watcher.Name(),
		),
	)
	assert.FatalOnError(t, e)
	compacting := s.NewSession()
	compacting.Announce(compacting.Name(), "building")
	compacting.CheckLive()
	assert.FatalOnError(
		t,
		s.Client().PostCompaction(
			compacting.Context,
			client.PostCompactionJSONRequestBody{
				Session: compacting.UUID,
				Phase:   client.Started,
			},
		),
	)
	notifications := delivery_tester.ClientKindEntries(
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
	compacting := s.NewSession()
	compacting.CheckLive()
	e := s.Client().PostCompaction(
		compacting.Context,
		client.PostCompactionJSONRequestBody{
			Session: compacting.UUID,
			Phase:   "halfway",
		},
	)
	assert.True(t, validation.Is(e))
}
