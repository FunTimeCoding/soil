package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/unit/markup_tester"
	"strings"
	"testing"
)

func TestKeywordSearchFollowsTypingAndMarksTerms(t *testing.T) {
	s, c := servedWeb(t)
	pushDocument(
		t,
		c,
		"pushed",
		"alfa.md",
		"# Alfa\n\nthe pipeline keeps a reranker warm",
	)
	_, page := servedPage(t, s, "/search?query=pipeline&mode=keyword")
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
	fragment := servedExtended(t, s, "/search?query=pipeline&mode=keyword")
	assert.True(t, !strings.Contains(fragment, "<html"))
	assert.StringContains(
		t,
		`<mark class="highlight">pipeline</mark>`,
		fragment,
	)
}

func TestHybridSearchKeepsItsButton(t *testing.T) {
	s, _ := servedWeb(t)
	_, page := servedPage(t, s, "/search")
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
