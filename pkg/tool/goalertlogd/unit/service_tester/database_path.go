package service_tester

import (
	"path/filepath"
	"testing"
)

func DatabasePath(t *testing.T) string {
	t.Helper()

	return filepath.Join(t.TempDir(), "goalertlogd.db")
}
