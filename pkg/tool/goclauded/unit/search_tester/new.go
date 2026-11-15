package search_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"path/filepath"
	"testing"
)

func New(t *testing.T) (*search_index.Index, string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "s1.jsonl")
	system.WriteFile(path, []byte(Transcript()), 0644)
	d := connection.NewMemory()
	t.Cleanup(func() { errors.PanicClose(d) })
	x := search_index.New(d)
	x.Append("s1", path)

	return x, path
}
