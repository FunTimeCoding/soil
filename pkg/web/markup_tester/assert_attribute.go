package markup_tester

func (o *Tester) AssertAttribute(
	expected string,
	body string,
	selector string,
	name string,
) {
	o.t.Helper()
	s := document(body).FindMatcher(o.matcher(selector)).First()

	if s.Length() == 0 {
		o.t.Errorf("%s: no match", selector)

		return
	}

	if actual, okay := s.Attr(name); !okay || actual != expected {
		o.t.Errorf(
			"%s [%s]: expected %q, got %q (present %t)",
			selector,
			name,
			expected,
			actual,
			okay,
		)
	}
}
