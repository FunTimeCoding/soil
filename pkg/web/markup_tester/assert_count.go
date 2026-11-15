package markup_tester

import "testing"

func AssertCount(
	t *testing.T,
	expected int,
	body string,
	selector string,
) {
	t.Helper()
	actual := document(body).FindMatcher(matcher(t, selector)).Length()

	if actual != expected {
		t.Errorf("%s: expected %d matches, got %d", selector, expected, actual)
	}
}
