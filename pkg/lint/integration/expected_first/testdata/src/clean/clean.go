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

type Tester struct {
	t *testing.T
}

func (o *Tester) AssertPath(
	expected string,
	path string,
) {
	o.t.Helper()

	if path != expected {
		o.t.Fatal(path)
	}
}
