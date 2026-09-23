package server

import (
	"fmt"
	"path/filepath"
	"testing"
)

func countingEntry(
	t *testing.T,
	name string,
) (string, string) {
	t.Helper()
	marker := filepath.Join(t.TempDir(), "runs")

	return marker, fmt.Sprintf(
		"%s: sh -c \"echo run >> %s; sleep 60\"\n",
		name,
		marker,
	)
}
