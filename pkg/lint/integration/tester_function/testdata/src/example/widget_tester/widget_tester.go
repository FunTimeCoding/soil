package widget_tester

import "testing"

type Tester struct {
	t *testing.T
}

func New(t *testing.T) *Tester {
	return &Tester{t: t}
}

func Build(t *testing.T) string {
	t.Helper()

	return "alfa"
}

func Seed(
	t *testing.T,
	name string,
) string {
	t.Helper()

	return name
}

func NewFixture(t *testing.T) string {
	t.Helper()

	return "bravo"
}

func Plain(name string) string {
	return name
}
