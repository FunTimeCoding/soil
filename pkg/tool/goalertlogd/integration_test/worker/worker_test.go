//go:build local

package worker

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/integration_test/worker_tester"
	"testing"
)

func TestPollSavesNewAlert(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic(
			"abc123",
			"HighMemory",
			"critical",
			"Memory above 90%",
		),
	)
	o.Worker.Poll()
	records := o.Store.MustByName("HighMemory")
	assert.Count(t, 1, records)
	assert.String(t, "abc123", records[0].Fingerprint)
	assert.String(t, "critical", records[0].Severity)
	assert.True(t, records[0].End == nil)
}

func TestPollResolvesRemovedAlert(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", "critical", ""),
	)
	o.Worker.Poll()
	o.MockClient.Remove("abc123")
	o.Worker.Poll()
	records := o.Store.MustByName("HighMemory")
	assert.Count(t, 1, records)
	assert.True(t, records[0].End != nil)
}

func TestPollIgnoresDuplicateFiring(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", "critical", ""),
	)
	o.Worker.Poll()
	o.Worker.Poll()
	o.Worker.Poll()
	assert.Count(t, 1, o.Store.MustByName("HighMemory"))
}

func TestRecoverStaleAdoptsFiring(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", "critical", ""),
	)
	o.Worker.Poll()
	fresh := worker_tester.NewWithStore(t, o.Store, o.MockClient)
	fresh.Worker.RecoverStale()
	assert.Count(t, 1, o.Store.MustUnresolved())
	fresh.Worker.Poll()
	assert.Count(t, 1, o.Store.MustByName("HighMemory"))
	assert.True(t, o.Store.MustByName("HighMemory")[0].End == nil)
}

func TestRecoverStaleResolvesGone(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", "critical", ""),
	)
	o.Worker.Poll()
	o.MockClient.Remove("abc123")
	fresh := worker_tester.NewWithStore(t, o.Store, o.MockClient)
	fresh.Worker.RecoverStale()
	assert.Count(t, 0, o.Store.MustUnresolved())
	records := o.Store.MustByName("HighMemory")
	assert.Count(t, 1, records)
	assert.True(t, records[0].End != nil)
}

func TestPollPrunesOldRecords(t *testing.T) {
	o := worker_tester.NewWithZeroRetention(t)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", "critical", ""),
	)
	o.Worker.Poll()
	o.MockClient.Remove("abc123")
	o.Worker.Poll()
	assert.Count(t, 0, o.Store.MustByName("HighMemory"))
}

func TestPollTracksLastPollTime(t *testing.T) {
	o := worker_tester.New(t)
	assert.True(t, o.Worker.LastPoll().IsZero())
	o.Worker.Poll()
	assert.False(t, o.Worker.LastPoll().IsZero())
}
