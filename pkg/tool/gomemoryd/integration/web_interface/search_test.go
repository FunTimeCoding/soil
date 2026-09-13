package web_interface

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
	"strings"
	"testing"
)

func body(
	t *testing.T,
	address string,
	extended bool,
) string {
	t.Helper()
	q, e := http.NewRequest(http.MethodGet, address, nil)
	assert.FatalOnError(t, e)

	if extended {
		q.Header.Set(constant.ExtendedRequest, "true")
	}

	r, f := http.DefaultClient.Do(q)
	assert.FatalOnError(t, f)
	defer errors.PanicClose(r.Body)
	b, g := io.ReadAll(r.Body)
	assert.FatalOnError(t, g)

	return string(b)
}

func TestSearchPartialWordMatches(t *testing.T) {
	s := base.New(t)
	defer s.Close()
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
	defer s.Close()
	address := fmt.Sprintf("http://localhost:%d/search?query=scr", s.Port)
	assert.True(t, !strings.Contains(body(t, address, true), "<html"))
	assert.StringContains(t, "<html", body(t, address, false))
}

func TestSearchPageCarriesLiveWiring(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	address := fmt.Sprintf("http://localhost:%d/search", s.Port)
	page := body(t, address, false)
	assert.StringContains(t, "keyup changed delay:200ms", page)
	assert.StringContains(t, "search-results", page)
	assert.StringContains(t, "search-indicator", page)
}
