//go:build local

package importer

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"testing"
)

func TestImportFromFixtures(t *testing.T) {
	s := store.New(connection.NewMemory())
	defer s.Close()
	result, e := Import(s, fixture.Path(constant.MemoryPath))

	if e != nil {
		t.Fatal(e)
	}

	if result.Created != 2 {
		t.Fatalf("expected 2 created, got %d", result.Created)
	}

	result2, e := Import(s, fixture.Path(constant.MemoryPath))

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
