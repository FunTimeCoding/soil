package unit

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func inputFile(
	t *testing.T,
	name string,
) string {
	t.Helper()
	result := filepath.Join(t.TempDir(), name)
	system.SaveFile(result, "log")

	return result
}
