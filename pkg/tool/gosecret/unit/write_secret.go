package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosecret/constant"
	"os"
	"path/filepath"
	"testing"
)

func writeSecret(t *testing.T) string {
	path := filepath.Join(t.TempDir(), "example-secret.yaml")
	errors.PanicOnError(os.WriteFile(path, []byte(constant.TestManifest), 0644))

	return path
}
