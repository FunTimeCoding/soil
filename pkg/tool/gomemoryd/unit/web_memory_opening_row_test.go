package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/markup_tester"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestMemoryRowOpensEditsAndSavesInPlace(t *testing.T) {
	s := base.New(t)
	o := save_option.New()
	o.Name = "alfa"
	o.Description = "first"
	o.Content = "alfa content"
	o.Type = "feedback"
	identifier, e := s.Store().CreateMemory(o, time.Now())
	assert.FatalOnError(t, e)
	page := s.Body(t, constant.MemoriesPath)
	markup_tester.AssertCount(t, 1, page, "tbody tr.opening-row")
	markup_tester.AssertAttribute(
		t,
		"/memories/1/open",
		page,
		"tr.opening-row",
		"hx-get",
	)
	opened := s.ExtendedBody(t, "/memories/1/open")
	assert.StringContains(t, `class="opening-detail"><td colspan="3">`, opened)
	markup_tester.AssertText(t, "alfa content", opened, "pre")
	markup_tester.AssertAttribute(
		t,
		"/memories/1/edit",
		opened,
		".action-row a[hx-get]",
		"href",
	)
	editor := s.ExtendedBody(t, "/memories/1/edit")
	markup_tester.AssertText(
		t,
		"alfa content",
		editor,
		"textarea[name=content]",
	)
	refused := s.PostBody(
		t,
		"/memories/1/edit",
		url.Values{
			constant.MemoryName:  {"alfa"},
			constant.Description: {"two\nlines"},
			constant.Content:     {"typed"},
		},
	)
	assert.StringContains(t, "description must be one line", refused)
	markup_tester.AssertText(t, "typed", refused, "textarea[name=content]")
	saved := s.PostBody(
		t,
		"/memories/1/edit",
		url.Values{
			constant.MemoryName:  {"alfa"},
			constant.Description: {"rewritten"},
			constant.Content:     {"alfa edited"},
		},
	)
	assert.StringContains(t, "Saved", saved)
	markup_tester.AssertText(t, "alfa edited", saved, "pre")
	markup_tester.AssertText(t, "rewritten", saved, "td.clamp-cell span")
	history, f := s.Store().GetMemoryHistory(identifier)
	assert.FatalOnError(t, f)
	assert.String(t, "web", history[len(history)-1].Source)
	closed := s.ExtendedBody(t, "/memories/1/close")
	assert.True(t, !strings.Contains(closed, "alfa edited"))
	markup_tester.AssertAttribute(
		t,
		"/memories/1/open",
		closed,
		"template tr.opening-row",
		"hx-get",
	)
}

func TestDocumentSourcedMemoryOpensReadOnly(t *testing.T) {
	s := base.New(t)
	o := save_option.New()
	o.Name = "lore"
	o.Content = "lore content"
	o.Type = "reference"
	o.ProvenanceFile = "doc/lore.md"
	identifier, e := s.Store().CreateMemory(o, time.Now())
	assert.FatalOnError(t, e)
	opened := s.ExtendedBody(t, "/memories/1/open")
	assert.StringContains(t, "loaded from doc/lore.md", opened)
	markup_tester.AssertCount(t, 0, opened, `a[href="/memories/1/edit"]`)
	refused := s.PostBody(
		t,
		"/memories/1/edit",
		url.Values{
			constant.MemoryName: {"lore"},
			constant.Content:    {"overwritten"},
		},
	)
	assert.StringContains(t, "document-sourced and read-only", refused)
	m, f := s.Store().GetMemory(identifier)
	assert.FatalOnError(t, f)
	assert.String(t, "lore content", m.Content)
}

func TestHiddenMemoryRefusesWebEdit(t *testing.T) {
	s := base.NewHidden(t, "private")
	o := save_option.New()
	o.Name = "zulu"
	o.Content = "zulu content"
	o.Type = "feedback"
	o.Tags = []string{"private"}
	identifier, e := s.Store().CreateMemory(o, time.Now())
	assert.FatalOnError(t, e)
	answer := s.PostBody(
		t,
		"/memories/1/edit",
		url.Values{
			constant.MemoryName: {"zulu"},
			constant.Content:    {"overwritten"},
		},
	)
	assert.StringContains(t, "memory not found: 1", answer)
	m, f := s.Store().GetMemory(identifier)
	assert.FatalOnError(t, f)
	assert.String(t, "zulu content", m.Content)
}
