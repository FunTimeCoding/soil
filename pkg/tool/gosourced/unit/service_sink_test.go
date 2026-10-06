package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"os"
	"path/filepath"
	"testing"
)

func TestSinkHoldsWritesUntilCommit(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "a.go")
	s := sink.New(root)
	assert.True(t, s.Empty())
	s.Write(path, []byte("first"))
	s.Write(path, []byte("second"))
	assert.False(t, s.Empty())
	assert.False(t, system.FileExists(path))
	assert.FatalOnError(t, s.Commit())
	content, e := os.ReadFile(path)
	assert.FatalOnError(t, e)
	assert.String(t, "second", string(content))
}

func TestSinkDropsOperationsOutsideRoot(t *testing.T) {
	base := t.TempDir()
	root := filepath.Join(base, "user")
	assert.FatalOnError(t, os.MkdirAll(root, 0755))
	inside := filepath.Join(root, "a.go")
	assert.FatalOnError(t, os.WriteFile(inside, []byte("user"), 0644))
	outside := filepath.Join(base, "library", "a.go")
	sibling := filepath.Join(base, "user-other", "a.go")
	s := sink.New(root)
	s.Write(outside, []byte("library"))
	s.Write(sibling, []byte("sibling"))
	s.Rename(inside, outside)
	s.Remove(sibling)
	assert.True(t, s.Empty())
	assert.Strings(t, []string{outside, sibling, outside, sibling}, s.Dropped())
	assert.FatalOnError(t, s.Commit())
	assert.False(t, system.FileExists(outside))
	assert.False(t, system.FileExists(sibling))
	assert.True(t, system.FileExists(inside))
}

func TestSinkCommitReplaysInOrder(t *testing.T) {
	root := t.TempDir()
	written := filepath.Join(root, "a.go")
	removed := filepath.Join(root, "b.go")
	assert.FatalOnError(t, os.WriteFile(removed, []byte("b"), 0644))
	moved := filepath.Join(root, "target", "a.go")
	s := sink.New(root)
	s.MakeDirectory(filepath.Join(root, "target"))
	s.Write(written, []byte("a"))
	s.Rename(written, moved)
	s.Remove(removed)
	assert.FatalOnError(t, s.Commit())
	content, e := os.ReadFile(moved)
	assert.FatalOnError(t, e)
	assert.String(t, "a", string(content))
	assert.False(t, system.FileExists(written))
	assert.False(t, system.FileExists(removed))
}
