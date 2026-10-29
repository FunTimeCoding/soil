package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/worker_tester"
	"testing"
)

func TestPollSavesNewAlert(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic(
			"abc123",
			"HighMemory",
			constant.CriticalSeverity,
			"Memory above 90%",
		),
	)
	o.Worker.Poll()
	records := o.Alerts("HighMemory")
	assert.Count(t, 1, records)
	assert.String(t, "abc123", records[0].Fingerprint)
	assert.String(t, "critical", records[0].Severity)
	assert.True(t, records[0].End == nil)
}

func TestPollResolvesRemovedAlert(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", constant.CriticalSeverity, ""),
	)
	o.Worker.Poll()
	o.MockClient.Remove("abc123")
	o.Worker.Poll()
	records := o.Alerts("HighMemory")
	assert.Count(t, 1, records)
	assert.True(t, records[0].End != nil)
}

func TestPollIgnoresDuplicateFiring(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", constant.CriticalSeverity, ""),
	)
	o.Worker.Poll()
	o.Worker.Poll()
	o.Worker.Poll()
	assert.Count(t, 1, o.Alerts("HighMemory"))
}

func TestPollIgnoresDuplicateInResponse(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", constant.CriticalSeverity, ""),
	)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", constant.CriticalSeverity, ""),
	)
	o.Worker.Poll()
	assert.Count(t, 1, o.Alerts("HighMemory"))
}

func TestRestartPollKeepsFiringOpen(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", constant.CriticalSeverity, ""),
	)
	o.Worker.Poll()
	fresh := o.Restart()
	fresh.Worker.Poll()
	assert.Integer(t, 1, fresh.Firing())
	records := fresh.Alerts("HighMemory")
	assert.Count(t, 1, records)
	assert.True(t, records[0].End == nil)
}

func TestRestartPollResolvesGone(t *testing.T) {
	o := worker_tester.New(t)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", constant.CriticalSeverity, ""),
	)
	o.Worker.Poll()
	o.MockClient.Remove("abc123")
	fresh := o.Restart()
	fresh.Worker.Poll()
	assert.Integer(t, 0, fresh.Firing())
	records := fresh.Alerts("HighMemory")
	assert.Count(t, 1, records)
	assert.True(t, records[0].End != nil)
}

func TestPollPrunesOldRecords(t *testing.T) {
	o := worker_tester.NewWith(
		t,
		service_tester.DatabasePath(t),
		service_tester.NewAlertmanager(t),
		0,
	)
	o.MockClient.Add(
		alert.NewBasic("abc123", "HighMemory", constant.CriticalSeverity, ""),
	)
	o.Worker.Poll()
	o.MockClient.Remove("abc123")
	o.Worker.Poll()
	assert.Count(t, 0, o.Alerts("HighMemory"))
}

func TestPollTracksLastPollTime(t *testing.T) {
	o := worker_tester.New(t)
	assert.True(t, o.Worker.LastPoll().IsZero())
	o.Worker.Poll()
	assert.False(t, o.Worker.LastPoll().IsZero())
}
