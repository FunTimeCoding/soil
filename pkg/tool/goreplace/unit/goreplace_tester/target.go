package goreplace_tester

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func Target(
	t *testing.T,
	content string,
) string {
	t.Helper()
	directory := t.TempDir()
	system_tester.New(t).Write(directory, "target.md", content)

	return filepath.Join(directory, "target.md")
}
