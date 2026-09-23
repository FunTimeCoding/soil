package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"testing"
)

func TestSendNotificationRefusesACallsignNobodyHolds(t *testing.T) {
	s := service_tester.New(t)
	_, e := s.Service.SendNotification(
		"Nobody",
		"downloader",
		"a file landed",
		false,
	)
	assert.True(t, e != nil)
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "Nobody", e.Error())
}

func TestSendNotificationReachesTheHolder(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	_, f := s.Service.SendNotification(
		r.Callsign,
		"downloader",
		"a file landed",
		false,
	)
	assert.FatalOnError(t, f)
	drained := s.Check("session-1")
	assert.Count(
		t,
		1,
		fixture.EntriesByKind(drained.Entries, constant.QueueNotification),
	)
}
