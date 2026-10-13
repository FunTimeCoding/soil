package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"strings"
	"testing"
)

func TestSearchPartialWordMatches(t *testing.T) {
	s := base.New(t)
	o := save_option.New()
	o.Name = "scripted rewrites"
	o.Content = "scripted rewrites bypass the read guard"
	o.Type = "feedback"
	_, e := s.Store().CreateMemory(o)
	assert.FatalOnError(t, e)
	target := "/search?query=scr"
	assert.StringContains(
		t,
		"scripted rewrites",
		body(t, s.Server, target, false),
	)
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
	assert.StringContains(t, "keyup changed delay:200ms", page)
	assert.StringContains(t, "search-results", page)
	assert.StringContains(t, "search-indicator", page)
}
