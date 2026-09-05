package assert

import "testing"

func String(
	t *testing.T,
	expected string,
	actual string,
) {
	t.Helper()

	if expected != actual {
		t.Fail()
	}
}

func Item(
	t *testing.T,
	expected any,
	actual any,
) {
	t.Helper()

	if expected != actual {
		t.Fail()
	}
}

func True(
	t *testing.T,
	actual bool,
) {
	t.Helper()

	if !actual {
		t.Fail()
	}
}
