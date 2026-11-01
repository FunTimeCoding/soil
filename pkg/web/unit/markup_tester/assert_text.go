package markup_tester

import (
	"strings"
	"testing"
)

func AssertText(
	t *testing.T,
	expected string,
	body string,
	selector string,
) {
	t.Helper()
	s := document(body).FindMatcher(matcher(t, selector)).First()

	if s.Length() == 0 {
		t.Errorf("%s: no match", selector)

		return
	}

	if actual := strings.TrimSpace(s.Text()); actual != expected {
		t.Errorf("%s: expected text %q, got %q", selector, expected, actual)
	}
}
