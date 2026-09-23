package flagged

import "testing"

func Exercise(t *testing.T) {
	assertValue(t, "actual", "expected")
	assertBetween(t, 1, 2, 3)
	assertShortSpelling(t, 1, true)
}

func assertShortSpelling(
	t *testing.T,
	value int,
	expect bool,
) {
	t.Helper()

	if (value > 0) != expect {
		t.Fatal(value)
	}
}

type Tester struct {
	t *testing.T
}

func (o *Tester) AssertPath(
	path string,
	expected string,
) {
	o.t.Helper()

	if path != expected {
		o.t.Fatal(path)
	}
}

func assertValue(
	t *testing.T,
	actual string,
	expected string,
) {
	t.Helper()

	if actual != expected {
		t.Fatal(actual)
	}
}

func assertBetween(
	t *testing.T,
	low int,
	expected int,
	high int,
) {
	t.Helper()

	if expected < low || expected > high {
		t.Fatal(expected)
	}
}
