package markup_tester

import (
	"github.com/andybalholm/cascadia"
	"testing"
)

func matcher(
	t *testing.T,
	selector string,
) cascadia.Selector {
	t.Helper()
	m, e := cascadia.Compile(selector)

	if e != nil {
		t.Fatalf("selector %q does not compile: %v", selector, e)
	}

	return m
}
