package clean

import "testing"

func Exercise(t *testing.T) {
	assertValue(t, "expected", "actual")
	assertContent(t, "content")
	checkValue(t, "actual", "expected")
}

func assertValue(
	t *testing.T,
	expected string,
	actual string,
) {
	t.Helper()

	if actual != expected {
		t.Fatal(actual)
	}
}

func assertContent(
	t *testing.T,
	content string,
) {
	t.Helper()

	if content == "" {
		t.Fatal(content)
	}
}

func checkValue(
	t *testing.T,
	actual string,
	expected string,
) {
	t.Helper()

	if actual != expected {
		t.Fatal(actual)
	}
}
