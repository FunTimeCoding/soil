package web_interface

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
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
	address := fmt.Sprintf("http://localhost:%d/search?query=scr", s.Port)
	assert.StringContains(t, "scripted rewrites", body(t, address, false))
}

func TestSearchFragmentOmitsPageChrome(t *testing.T) {
	s := base.New(t)
	address := fmt.Sprintf("http://localhost:%d/search?query=scr", s.Port)
	assert.True(t, !strings.Contains(body(t, address, true), "<html"))
	assert.StringContains(t, "<html", body(t, address, false))
}

func TestSearchPageCarriesLiveWiring(t *testing.T) {
	s := base.New(t)
	address := fmt.Sprintf("http://localhost:%d/search", s.Port)
	page := body(t, address, false)
	assert.StringContains(t, "keyup changed delay:200ms", page)
	assert.StringContains(t, "search-results", page)
	assert.StringContains(t, "search-indicator", page)
}
