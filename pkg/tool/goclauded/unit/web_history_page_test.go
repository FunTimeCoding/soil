package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/history_page"
	"github.com/funtimecoding/soil/pkg/web/unit/markup_tester"
	"net/http"
	"testing"
)

func TestHistoryKindChipsToggleTheirKind(t *testing.T) {
	var b bytes.Buffer
	assert.FatalOnError(
		t,
		history_page.Filters([]string{constant.Complete}).Render(&b),
	)
	assert.StringContains(
		t,
		`<a href="/history" class="chip active" aria-current="true">Completions</a>`,
		b.String(),
	)
	assert.StringContains(
		t,
		`<a href="/history?kind=complete&amp;kind=summarize" class="chip">Summaries</a>`,
		b.String(),
	)
}

func TestHistoryEditRowRoundTrips(t *testing.T) {
	s := base.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Moment("session-1", r.Callsign, "figured out the root cause")
	events := s.Store.Events(
		event_query.New().Kind(constant.Moment).SetLimit(1),
	)
	assert.Count(t, 1, events)
	edit := body(
		t,
		s.Server,
		history_page.EventPath(events[0].Identifier, true),
	)
	markup_tester.AssertCount(t, 1, edit, "textarea")
	assert.StringContains(t, "figured out the root cause", edit)
	assert.StringContains(t, "Cancel", edit)
	row := body(
		t,
		s.Server,
		history_page.EventPath(events[0].Identifier, false),
	)
	markup_tester.AssertCount(t, 0, row, "textarea")
	assert.StringContains(t, "figured out the root cause", row)
	assert.Integer(
		t,
		http.StatusNotFound,
		pageStatus(t, s, history_page.EventPath(99999, true)),
	)
}
