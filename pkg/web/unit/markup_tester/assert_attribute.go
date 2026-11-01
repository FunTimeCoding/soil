package markup_tester

import "testing"

func AssertAttribute(
	t *testing.T,
	expected string,
	body string,
	selector string,
	name string,
) {
	t.Helper()
	s := document(body).FindMatcher(matcher(t, selector)).First()

	if s.Length() == 0 {
		t.Errorf("%s: no match", selector)

		return
	}

	if actual, okay := s.Attr(name); !okay || actual != expected {
		t.Errorf(
			"%s [%s]: expected %q, got %q (present %t)",
			selector,
			name,
			expected,
			actual,
			okay,
		)
	}
}
