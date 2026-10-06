package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/index/cache"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
	"sync"
	"testing"
)

func TestCacheReusesAnUnchangedWorkspace(t *testing.T) {
	directory := kindModule(t)
	c := cache.New(t.TempDir())
	first := c.Workspace(directory)
	assert.True(t, first == c.Workspace(directory))
}

func TestCacheRebuildsAfterAnEdit(t *testing.T) {
	directory := kindModule(t)
	c := cache.New(t.TempDir())
	first := c.Workspace(directory)
	testutil.WriteFile(
		t,
		directory,
		"two/more.go",
		"package two\n\nfunc More() string {\n\treturn \"more\"\n}\n",
	)
	assert.False(t, first == c.Workspace(directory))
}

func TestCacheRebuildsAfterAnEditInAReplacedModule(t *testing.T) {
	library, user := replacedModules(t)
	c := cache.New(t.TempDir())
	first := c.Workspace(user)
	assert.True(t, first == c.Workspace(user))
	testutil.WriteFile(
		t,
		library,
		"extra.go",
		"package lib\n\nfunc Extra() string {\n\treturn \"extra\"\n}\n",
	)
	assert.False(t, first == c.Workspace(user))
}

func TestSnapshotSame(t *testing.T) {
	directory := kindModule(t)
	before := snapshot.Take(directory)
	assert.True(t, before.Same(snapshot.Take(directory)))
	testutil.WriteFile(t, directory, "one/one.go", "package one\n")
	assert.False(t, before.Same(snapshot.Take(directory)))
}

func TestWorkspaceReferencesAreSafeInParallel(t *testing.T) {
	w := index.New(t.TempDir(), kindModule(t))
	var group sync.WaitGroup
	results := make([]any, 8)

	for i := range results {
		group.Go(
			func() {
				results[i] = w.References()
			},
		)
	}

	group.Wait()

	for _, r := range results {
		assert.True(t, r == results[0])
	}
}
