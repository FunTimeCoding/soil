package plain

import "testing"

func Build(t *testing.T) string {
	t.Helper()

	return "charlie"
}

func Seed(t *testing.T) string {
	t.Helper()

	return "delta"
}
