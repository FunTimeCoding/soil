package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"testing"
)

func TestAbandonRunningFailsOnlyTheRunsLeftRunning(t *testing.T) {
	m := lite.NewMemory()
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m, "runs")
	s := store.New(m, "runs")
	left := s.NewRun()
	left.Status = constant.StoreStatusRunning
	s.Create(left)
	done := s.NewRun()
	done.Status = constant.StoreStatusSuccess
	done.Output = "applied"
	s.Create(done)
	s.AbandonRunning()
	abandoned, e := s.Find(left.Identifier)
	assert.FatalOnError(t, e)
	assert.String(t, "error", abandoned.Status)
	assert.String(
		t,
		"abandoned: the daemon stopped before the run finished",
		abandoned.ErrorOutput,
	)
	kept, f := s.Find(done.Identifier)
	assert.FatalOnError(t, f)
	assert.String(t, "success", kept.Status)
	assert.String(t, "", kept.ErrorOutput)
	assert.String(t, "applied", kept.Output)
}
