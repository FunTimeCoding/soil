package unit

import (
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"path/filepath"
	"testing"
)

func indexedTranscript(t *testing.T) (*search_index.Index, string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "s1.jsonl")
	system.WriteFile(path, []byte(searchTranscript()), 0644)
	x := search_index.New(connection.NewMemory())
	x.Append("s1", path)

	return x, path
}
