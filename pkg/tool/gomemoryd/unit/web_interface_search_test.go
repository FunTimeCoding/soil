package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"strings"
	"testing"
	"time"
)

func TestSearchPartialWordMatches(t *testing.T) {
	s := base.New(t)
	o := save_option.New()
	o.Name = "scripted rewrites"
	o.Content = "scripted rewrites bypass the read guard"
	o.Type = "feedback"
	_, e := s.Store().CreateMemory(o, time.Now())
	assert.FatalOnError(t, e)
	target := "/search?query=scr"
	assert.StringContains(
		t,
		"scripted rewrites",
		body(t, s.Server, target, false),
	)
}

func TestSearchOmitsHiddenMemories(t *testing.T) {
	s := base.NewHidden(t, "private")
	visible := save_option.New()
	visible.Name = "scripted reads"
	visible.Content = "scripted reads stay visible"
	visible.Type = "feedback"
	_, e := s.Store().CreateMemory(visible, time.Now())
	assert.FatalOnError(t, e)
	hidden := save_option.New()
	hidden.Name = "scripted rewrites"
	hidden.Content = "scripted rewrites stay hidden"
	hidden.Type = "feedback"
	hidden.Tags = []string{"private"}
	_, f := s.Store().CreateMemory(hidden, time.Now())
	assert.FatalOnError(t, f)
	target := "/search?query=scr"

	for _, extended := range []bool{false, true} {
		page := body(t, s.Server, target, extended)
		assert.StringContains(t, "scripted reads", page)
		assert.True(t, !strings.Contains(page, "scripted rewrites"))
		assert.True(t, !strings.Contains(page, "private"))
	}
}

func TestSearchFragmentOmitsPageChrome(t *testing.T) {
	s := base.New(t)
	target := "/search?query=scr"
	assert.True(t, !strings.Contains(body(t, s.Server, target, true), "<html"))
	assert.StringContains(t, "<html", body(t, s.Server, target, false))
}

func TestSearchPageCarriesLiveWiring(t *testing.T) {
	s := base.New(t)
	target := "/search"
	page := body(t, s.Server, target, false)
	assert.StringContains(t, `hx-trigger="input delay:200ms, submit"`, page)
	assert.StringContains(t, `hx-target="#search-results"`, page)
	assert.StringContains(t, `hx-push-url="true"`, page)
	assert.StringContains(t, `<div id="search-results">`, page)
}

func TestSearchMarksStemmedMatchInSnippet(t *testing.T) {
	s := base.New(t)
	o := save_option.New()
	o.Name = "boathouse"
	o.Content = "the crew kept rowing home before dark"
	o.Type = "feedback"
	_, e := s.Store().CreateMemory(o, time.Now())
	assert.FatalOnError(t, e)
	assert.StringContains(
		t,
		`kept <mark class="highlight">rowing</mark> home`,
		body(t, s.Server, "/search?query=rowed", true),
	)
}
