package flagged

import "testing"

func Exercise(t *testing.T) {
	assertValue(t, "actual", "expected")
	assertBetween(t, 1, 2, 3)
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
