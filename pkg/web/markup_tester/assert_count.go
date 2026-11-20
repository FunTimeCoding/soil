package markup_tester

func (o *Tester) AssertCount(
	expected int,
	body string,
	selector string,
) {
	o.t.Helper()
	actual := document(body).FindMatcher(o.matcher(selector)).Length()

	if actual != expected {
		o.t.Errorf(
			"%s: expected %d matches, got %d",
			selector,
			expected,
			actual,
		)
	}
}
