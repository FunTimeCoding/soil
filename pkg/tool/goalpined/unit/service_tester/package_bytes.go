package service_tester

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/unit/package_server_tester"
	"path/filepath"
	"testing"
)

func PackageBytes(t *testing.T) []byte {
	t.Helper()
	directory := t.TempDir()
	path := filepath.Join(directory, "test.apk")
	package_server_tester.CreateTestPackage(directory, path)

	return system.ReadBytesUnsafe(path)
}
