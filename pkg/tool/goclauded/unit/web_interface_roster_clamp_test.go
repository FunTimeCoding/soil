package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/argument/edit_session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/web/unit/markup_tester"
	"testing"
)

func TestRosterCardClampsDescriptionAndKeepsItWhole(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("session-1")
	assert.FatalOnError(
		t,
		s.Store.Store.EditSession(
			"session-1",
			edit_session.New().WithDescription(
				"Charted the session from start to end, every turn read and every file named",
			),
		),
	)
	page := body(t, s.Server, constant.DashboardPath)
	markup_tester.AssertCount(t, 1, page, "p.card-meta.card-clamp")
	assert.StringContains(
		t,
		"Charted the session from start to end, every turn read and every file named",
		page,
	)
}

func TestRosterCardClampsTopicAndKeepsItWhole(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("session-1")
	assert.FatalOnError(
		t,
		s.Store.Store.EditSession(
			"session-1",
			edit_session.New().WithTopic(
				"lead: lanes for the language, the census, the wire and the whiteboard",
			),
		),
	)
	page := body(t, s.Server, constant.DashboardPath)
	markup_tester.AssertCount(t, 1, page, "p.card-subtitle.card-clamp")
	assert.StringContains(
		t,
		"lead: lanes for the language, the census, the wire and the whiteboard",
		page,
	)
}
