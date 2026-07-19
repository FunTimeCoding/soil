package unit_test

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteRead(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("a/b.yaml", "content")
	assert.String(t, "content", v.ReadString("a/b.yaml"))
}

func TestWriteReadBytes(t *testing.T) {
	v := virtual_file_system.New()
	v.Write("a.bin", []byte{1, 2, 3})
	assert.Any(t, []byte{1, 2, 3}, v.Read("a.bin"))
}

func TestHas(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("a.yaml", "x")
	assert.Boolean(t, true, v.Has("a.yaml"))
	assert.Boolean(t, false, v.Has("b.yaml"))
}

func TestDelete(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("a.yaml", "x")
	v.Delete("a.yaml")
	assert.Boolean(t, false, v.Has("a.yaml"))
}

func TestDeleteFlush(t *testing.T) {
	d := t.TempDir()
	v := virtual_file_system.New()
	v.WriteString("a.txt", "hello")
	v.Flush(d)
	v2 := virtual_file_system.New()
	v2.Delete("a.txt")
	v2.Flush(d)
	_, e := os.Stat(filepath.Join(d, "a.txt"))
	assert.True(t, os.IsNotExist(e))
}

func TestFilesSorted(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("b.yaml", "")
	v.WriteString("a.yaml", "")
	v.WriteString("c.yaml", "")
	assert.Any(t, []string{"a.yaml", "b.yaml", "c.yaml"}, v.Files())
}

func TestVirtualFileSystemDirectoryExists(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotest/server/router.go", "")
	v.WriteString("pkg/tool/gotest/constant/constant.go", "")
	assert.Boolean(t, true, v.DirectoryExists("pkg/tool/gotest"))
	assert.Boolean(
		t,
		true,
		v.DirectoryExists("pkg/tool/gotest/server"),
	)
	assert.Boolean(t, true, v.DirectoryExists("pkg"))
	assert.Boolean(
		t,
		false,
		v.DirectoryExists("pkg/tool/gotest/model_context"),
	)
	assert.Boolean(t, false, v.DirectoryExists("pkg/tool/other"))
}

func TestReadDirectory(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/tool/gotest/server/router.go", "")
	v.WriteString("pkg/tool/gotest/server/handler.go", "")
	v.WriteString("pkg/tool/gotest/constant/constant.go", "")
	v.WriteString("pkg/tool/gotest/run.go", "")
	assert.Any(
		t,
		[]string{"constant", "run.go", "server"},
		v.MustReadDirectory("pkg/tool/gotest"),
	)
	assert.Any(
		t,
		[]string{"handler.go", "router.go"},
		v.MustReadDirectory("pkg/tool/gotest/server"),
	)
}

func TestReadDirectoryMissing(t *testing.T) {
	v := virtual_file_system.New()
	result, e := v.ReadDirectory("pkg/tool/missing")
	assert.True(t, result == nil)
	assert.NotNil(t, e)
	assert.True(t, errors.Is(e, virtual_file_system.ErrorDirectoryNotFound))
}

func TestFlushFrom(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("sub/file.txt", "hello")
	d := t.TempDir()
	v.Flush(d)
	back := virtual_file_system.From(d)
	assert.String(t, "hello", back.ReadString("sub/file.txt"))
}

func TestMoveFlush(t *testing.T) {
	d := t.TempDir()
	v := virtual_file_system.New()
	v.WriteString("root.txt", "content")
	v.Flush(d)
	v2 := virtual_file_system.New()
	v2.Move("root.txt", "sub/moved.txt")
	v2.Flush(d)
	_, e := os.Stat(filepath.Join(d, "root.txt"))
	assert.True(t, os.IsNotExist(e))
	back := virtual_file_system.From(d)
	assert.String(t, "content", back.ReadString("sub/moved.txt"))
}

func TestFromMetadata(t *testing.T) {
	d := t.TempDir()
	v := virtual_file_system.New()
	v.WriteString("a.txt", "hello")
	v.Flush(d)
	m := virtual_file_system.FromMetadata(d)
	assert.Boolean(t, true, m.Has("a.txt"))
	assert.True(t, m.Read("a.txt") == nil)
	f := m.FileAt("a.txt")
	assert.True(t, f.Size == 5)
	assert.Boolean(t, false, f.Loaded)
}
