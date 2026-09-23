package widget_tester

import "testing"

type Tester struct {
	t *testing.T
}

func New(t *testing.T) *Tester {
	return &Tester{t: t}
}

func (o *Tester) AssertName(
	expected string,
	actual string,
) {
	o.t.Helper()

	if expected != actual {
		o.t.Fail()
	}
}

func (o *Tester) Submit(
	t *testing.T,
	name string,
) string {
	t.Helper()

	return name
}
