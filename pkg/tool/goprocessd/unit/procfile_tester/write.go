package procfile_tester

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"path/filepath"
	"testing"
)

func Write(
	t *testing.T,
	content string,
) string {
	t.Helper()
	directory := t.TempDir()
	system_tester.New(t).Write(directory, "Procfile", content)

	return filepath.Join(directory, "Procfile")
}
