package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	goqueryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"os"
	"path/filepath"
	"testing"
)

func openTestStore(t *testing.T) *store.Store {
	t.Helper()

	return store.New(connection.NewMemory())
}

func writeFixture(
	t *testing.T,
	directory string,
	name string,
	content string,
) {
	t.Helper()
	path := filepath.Join(directory, name)
	errors.PanicOnError(os.MkdirAll(filepath.Dir(path), 0o755))
	errors.PanicOnError(os.WriteFile(path, []byte(content), 0o644))
}

func indexedTestStore(t *testing.T) *store.Store {
	t.Helper()
	s := openTestStore(t)
	s.AddCollection(
		"test",
		fixture.Path(system.SearchPath),
		goqueryd.DefaultGlob,
	)
	s.Index("test")

	return s
}
