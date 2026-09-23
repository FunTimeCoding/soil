package widget_tester

import "testing"

type Tester struct {
	t *testing.T
}

func New(t *testing.T) *Tester {
	return &Tester{t: t}
}

func (o *Tester) Good(name string) string {
	o.t.Helper()

	return name
}

func (o *Tester) Bad(
	t *testing.T,
	name string,
) string {
	t.Helper()

	return name
}

func (o *Tester) AlsoBad(t *testing.T) {
	t.Helper()
}
