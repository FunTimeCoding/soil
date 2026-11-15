package secret_tester

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"path/filepath"
	"testing"
)

func Manifest(t *testing.T) string {
	t.Helper()
	directory := t.TempDir()
	system_tester.Write(
		t,
		directory,
		"example-secret.yaml",
		constant.TestManifest,
	)

	return filepath.Join(directory, "example-secret.yaml")
}
