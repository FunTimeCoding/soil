package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"testing"
	"time"
)

func TestDeleteEmptySessionNeedsNoConfirmation(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("husk")
	r, e := s.Service.DeleteSession("husk", "")
	assert.FatalOnError(t, e)
	assert.String(t, "husk", r.Identifier)
	assert.True(t, s.Store.GetSession("husk") == nil)
}

func TestDeleteRefusesWithoutConfirmation(t *testing.T) {
	s := service_tester.New(t)
	writeSessionFile(s.Harbor, "doomed", "some-slug")
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	_, e := s.Service.DeleteSession("doomed", "")
	assert.True(t, e != nil)
	assert.String(
		t,
		"session has a transcript, confirmation required",
		e.Error(),
	)
	assert.True(t, s.Store.GetSession("doomed") != nil)
}

func TestDeleteRejectsStaleConfirmation(t *testing.T) {
	s := service_tester.New(t)
	writeSessionFile(s.Harbor, "doomed", "some-slug")
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	stale := s.Service.DeleteHash("doomed")
	s.Store.EnsureSession("doomed")
	_, e := s.Service.DeleteSession("doomed", stale)
	assert.True(t, e != nil)
	assert.String(t, "confirmation does not match this session", e.Error())
	assert.True(t, s.Store.GetSession("doomed") != nil)
}

func TestDeleteRejectsForeignConfirmation(t *testing.T) {
	s := service_tester.New(t)
	writeSessionFile(s.Harbor, "doomed", "some-slug")
	writeSessionFile(s.Harbor, "bystander", "other-slug")
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	_, e := s.Service.DeleteSession("doomed", s.Service.DeleteHash("bystander"))
	assert.True(t, e != nil)
	assert.String(t, "confirmation does not match this session", e.Error())
	assert.True(t, s.Store.GetSession("doomed") != nil)
}

func TestDeleteReceiptReportsWhatWentAway(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("rich")
	assert.FatalOnError(
		t,
		s.Store.Store.LogEvent(
			"rich",
			constant.Announce,
			r.Callsign,
			map[string]string{constant.Topic: "counted"},
		),
	)
	_, e := s.Store.Store.UpsertCompletion(
		"rich",
		r.Callsign,
		constant.Complete,
		"counted",
		"summary body",
	)
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		s.Store.Store.UpsertSummary("rich", r.Callsign, "summary body"),
	)
	_, f := s.Store.Store.SetLabel("rich", "role", "reviewer")
	assert.FatalOnError(t, f)
	assert.FatalOnError(
		t,
		s.Store.Store.SendPulse("rich", r.Callsign, "a pulse"),
	)
	result, g := s.Service.DeleteSession("rich", s.Service.DeleteHash("rich"))
	assert.FatalOnError(t, g)
	assert.Integer(t, 1, result.Events)
	assert.Integer(t, 1, result.EventMetadata)
	assert.Integer(t, 1, result.Completions)
	assert.Integer(t, 1, result.Summaries)
	assert.Integer(t, 1, result.Labels)
	assert.Integer(t, 1, result.Pulses)
	assert.True(t, s.Store.GetSession("rich") == nil)
}

func TestDeleteClearsCallsignQueueWindow(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("husk")
	assert.FatalOnError(
		t,
		s.Store.Store.PushQueue(
			"husk",
			r.Callsign,
			constant.QueueTimeout,
			"for husk",
		),
	)
	assert.FatalOnError(
		t,
		s.Store.Store.SendNotification("husk", r.Callsign, "Dale", "for husk"),
	)
	s.Store.Advance(time.Minute)
	result, e := s.Service.DeleteSession("husk", "")
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, result.Queue)
	assert.Integer(t, 1, result.Notifications)
}
