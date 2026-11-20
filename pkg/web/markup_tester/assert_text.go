package markup_tester

import "strings"

func (o *Tester) AssertText(
	expected string,
	body string,
	selector string,
) {
	o.t.Helper()
	s := document(body).FindMatcher(o.matcher(selector)).First()

	if s.Length() == 0 {
		o.t.Errorf("%s: no match", selector)

		return
	}

	if actual := strings.TrimSpace(s.Text()); actual != expected {
		o.t.Errorf("%s: expected text %q, got %q", selector, expected, actual)
	}
}
