//go:build local

package importer

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/constant"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"path/filepath"
	"testing"
)

func TestImportFromFixtures(t *testing.T) {
	s := store.New(filepath.Join(t.TempDir(), constant.TestDatabase))
	defer s.Close()
	result, e := Import(s, fixture.Path(system.MemoryPath))

	if e != nil {
		t.Fatal(e)
	}

	if result.Created != 2 {
		t.Fatalf("expected 2 created, got %d", result.Created)
	}

	result2, e := Import(s, fixture.Path(system.MemoryPath))

	if e != nil {
		t.Fatal(e)
	}

	if result2.Created != 0 {
		t.Fatalf("expected 0 created on re-import, got %d", result2.Created)
	}

	if result2.Skipped != 2 {
		t.Fatalf("expected 2 skipped on re-import, got %d", result2.Skipped)
	}
}
