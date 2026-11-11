package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/unit/markup_tester"
	"net/url"
	"testing"
	"time"
)

func TestRosterCardCountsWhatIsQueued(t *testing.T) {
	s := base.New(t)
	r := s.Store.EnsureSession("session-1")

	for _, text := range []string{"one", "two", "three"} {
		assert.FatalOnError(
			t,
			s.Store.Store.PushQueue(
				"session-1",
				r.Callsign,
				constant.QueueMessage,
				text,
			),
		)
	}

	assert.StringContains(
		t,
		"3 queued",
		body(t, s.Server, constant.DashboardPath),
	)
}

func TestRosterCardShowsNothingQueuedWhenNothingWaits(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("session-1")
	assert.StringNotContains(
		t,
		"queued",
		body(t, s.Server, constant.DashboardPath),
	)
}

func TestRosterCruiseSitsInFooterOutsideLink(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("session-1")
	page := body(t, s.Server, constant.DashboardPath)
	action := session_page.ActionPath("session-1", constant.CruiseAction)
	markup_tester.AssertCount(
		t,
		1,
		page,
		join.Empty(`article[hx-preserve] footer form[action="`, action, `"]`),
	)
	markup_tester.AssertCount(t, 0, page, "a form")
	markup_tester.AssertCount(t, 0, page, `[sse-swap^="roster-"] form`)
}

func TestRefusedGraceAnswersCardsFormWithItsAlert(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("session-1")
	answer := submit(
		t,
		s,
		session_page.ActionPath("session-1", constant.CruiseAction),
		url.Values{
			constant.GraceKind: {constant.CruiseMinutes},
			constant.Grace:     {"61"},
		},
		true,
	)
	markup_tester.AssertCount(
		t,
		1,
		answer,
		join.Empty(
			"form#",
			session_page.Event("session-1", constant.RosterCruisePart),
		),
	)
	assert.StringContains(t, "from 1 to 59, not &#34;61&#34;", answer)
	setting, e := s.Store.Store.CruiseSetting("session-1")
	assert.FatalOnError(t, e)
	assert.Nil(t, setting.Grace)
}

func TestDetailPageTakesGraceInRange(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("session-1")
	submit(
		t,
		s,
		session_page.ActionPath("session-1", constant.CruiseAction),
		url.Values{
			constant.GraceKind: {constant.CruiseMinutes},
			constant.Grace:     {"12"},
		},
		false,
	)
	setting, e := s.Store.Store.CruiseSetting("session-1")
	assert.FatalOnError(t, e)
	assert.Duration(t, 12*time.Minute, *setting.Grace)
}

func TestDetailPageTurnsCruiseOff(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("session-1")
	submit(
		t,
		s,
		session_page.ActionPath("session-1", constant.CruiseAction),
		url.Values{constant.GraceKind: {constant.CruiseOff}},
		false,
	)
	setting, e := s.Store.Store.CruiseSetting("session-1")
	assert.FatalOnError(t, e)
	assert.Duration(t, 0, *setting.Grace)
}
