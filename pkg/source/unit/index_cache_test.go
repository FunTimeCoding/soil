package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/index/cache"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"github.com/funtimecoding/soil/pkg/source/unit/kind_tester"
	"github.com/funtimecoding/soil/pkg/source/unit/module_graph_tester"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func TestCacheReusesUnchangedWorkspace(t *testing.T) {
	directory := kind_tester.New(t).Module()
	c := cache.New(t.TempDir())
	first := c.Workspace(directory)
	assert.True(t, first == c.Workspace(directory))
}

func TestCacheRebuildsAfterEdit(t *testing.T) {
	directory := kind_tester.New(t).Module()
	c := cache.New(t.TempDir())
	first := c.Workspace(directory)
	system_tester.New(t).Write(
		directory,
		"two/more.go",
		"package two\n\nfunc More() string {\n\treturn \"more\"\n}\n",
	)
	assert.False(t, first == c.Workspace(directory))
}

func TestCacheRebuildsAfterEditInReplacedModule(t *testing.T) {
	library, user := module_graph_tester.New(t).Replaced()
	c := cache.New(t.TempDir())
	first := c.Workspace(user)
	assert.True(t, first == c.Workspace(user))
	system_tester.New(t).Write(
		library,
		"extra.go",
		"package lib\n\nfunc Extra() string {\n\treturn \"extra\"\n}\n",
	)
	assert.False(t, first == c.Workspace(user))
}

func TestSnapshotSame(t *testing.T) {
	directory := kind_tester.New(t).Module()
	before := snapshot.Take(directory)
	assert.True(t, before.Same(snapshot.Take(directory)))
	system_tester.New(t).Write(directory, "one/one.go", "package one\n")
	assert.False(t, before.Same(snapshot.Take(directory)))
}
