package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/session_page"
	"github.com/funtimecoding/soil/pkg/web/markup_tester"
	"net/http"
	"testing"
)

func TestPulseFormSitsOutsideLiveRegion(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("session-1")
	page := s.Body(t, session_page.Path("session-1"))
	markup_tester.New(t).AssertCount(
		1,
		page,
		`form[action="/sessions/session-1/pulse"]`,
	)
	markup_tester.New(t).AssertCount(0, page, "[sse-swap] form")
}

func TestDashboardActivityIsReadOnly(t *testing.T) {
	s := base.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Moment("session-1", r.Callsign, "figured out the root cause")
	page := s.Body(t, constant.DashboardPath)
	assert.StringContains(t, "figured out the root cause", page)
	markup_tester.New(t).AssertCount(0, page, `a[href$="/edit"]`)
}

func TestUnknownSessionIsNotFound(t *testing.T) {
	s := base.New(t)
	assert.Integer(
		t,
		http.StatusNotFound,
		s.Status(t, session_page.Path("no-such-session")),
	)
}
