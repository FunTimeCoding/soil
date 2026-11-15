package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/document_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/markup_tester"
	"strings"
	"testing"
)

func TestKeywordSearchFollowsTypingAndMarksTerms(t *testing.T) {
	s := base.Serve(
		t,
		service_tester.NewOption(
			t,
			service_tester.NewEmbedder(t),
			service_tester.NewReranker(t),
		),
		memory.New(),
	)
	c := client_tester.Rest(t, s)
	document_tester.Push(
		t,
		c,
		"pushed",
		"alfa.md",
		"# Alfa\n\nthe pipeline keeps a reranker warm",
	)
	page := s.Body(t, "/search?query=pipeline&mode=keyword")
	markup_tester.AssertAttribute(t, "true", page, "form.entry", "hx-push-url")
	markup_tester.AssertAttribute(
		t,
		"keyword",
		page,
		"form.entry input[name=mode]",
		"value",
	)
	markup_tester.AssertText(t, "keyword", page, "a.chip[aria-current]")
	markup_tester.AssertText(
		t,
		"pipeline",
		page,
		"#search-results td mark.highlight",
	)
	fragment := s.ExtendedBody(t, "/search?query=pipeline&mode=keyword")
	assert.True(t, !strings.Contains(fragment, "<html"))
	assert.StringContains(
		t,
		`<mark class="highlight">pipeline</mark>`,
		fragment,
	)
}

func TestHybridSearchKeepsItsButton(t *testing.T) {
	s := base.Serve(
		t,
		service_tester.NewOption(
			t,
			service_tester.NewEmbedder(t),
			service_tester.NewReranker(t),
		),
		memory.New(),
	)
	page := s.Body(t, "/search")
	markup_tester.AssertCount(t, 0, page, "form.entry[hx-push-url]")
	markup_tester.AssertCount(t, 1, page, "form.entry button[type=submit]")
	markup_tester.AssertText(t, "hybrid", page, "a.chip[aria-current]")
	markup_tester.AssertAttribute(
		t,
		"/search?mode=keyword&query=",
		page,
		"a.chip:not([aria-current])",
		"href",
	)
}
