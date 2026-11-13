package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"testing"
)

func TestLayerReadsThroughParentToDisk(t *testing.T) {
	root := layerModule(t)
	base := layer.New(nil, root)
	base.Write(filepath.Join(root, "alfa/alfa.go"), []byte("package alfa\n\n"))
	child := layer.New(base)
	written, e := child.Read(filepath.Join(root, "alfa/alfa.go"))
	assert.FatalOnError(t, e)
	disk, f := child.Read(filepath.Join(root, "charlie/charlie.go"))
	assert.FatalOnError(t, f)
	assert.String(t, "package alfa\n\n", string(written))
	assert.String(t, "package charlie\n", string(disk))
	assert.Integer(t, 1, child.Writes())
}

func TestLayerMoveShowsFileOnlyAtTarget(t *testing.T) {
	root := layerModule(t)
	l := layer.New(nil, root)
	from := filepath.Join(root, "alfa/bravo.go")
	to := filepath.Join(root, "alfa/foxtrot.go")
	l.Move(from, to)
	moved, e := l.Read(to)
	assert.FatalOnError(t, e)
	assert.String(t, "package alfa\n\nfunc Bravo() {}\n", string(moved))
	assert.False(t, l.Exists(from))
}

func TestLayerOverlayStubsRemovedGoFile(t *testing.T) {
	root := layerModule(t)
	l := layer.New(nil, root)
	l.Move(
		filepath.Join(root, "alfa/bravo.go"),
		filepath.Join(root, "alfa/foxtrot.go"),
	)
	overlay := l.Overlay()
	assert.String(
		t,
		"//go:build ignore\n\npackage alfa\n",
		string(overlay[filepath.Join(root, "alfa/bravo.go")]),
	)
	assert.String(
		t,
		"package alfa\n\nfunc Bravo() {}\n",
		string(overlay[filepath.Join(root, "alfa/foxtrot.go")]),
	)
}

func TestLayerChangesListMoveOnceWithDestination(t *testing.T) {
	root := layerModule(t)
	l := layer.New(nil, root)
	from := filepath.Join(root, "alfa/bravo.go")
	to := filepath.Join(root, "alfa/foxtrot.go")
	l.Move(from, to)
	changes := l.Changes()
	assert.Count(t, 1, changes)
	assert.String(t, from, changes[0].Path)
	assert.String(t, to, changes[0].Destination)
}

func TestLayerCommitWritesAndRemoves(t *testing.T) {
	root := layerModule(t)
	l := layer.New(nil, root)
	from := filepath.Join(root, "alfa/bravo.go")
	to := filepath.Join(root, "alfa/foxtrot.go")
	l.Move(from, to)
	refused, e := l.Commit(nil)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, refused)
	assert.False(t, system.FileExists(from))
	assert.True(t, system.FileExists(to))
}

func TestLayerCommitRefusesConcurrentWriteWithoutGraph(t *testing.T) {
	root := layerModule(t)
	l := layer.New(nil, root)
	l.Write(filepath.Join(root, "alfa/alfa.go"), []byte("package alfa\n\n"))
	charlie := filepath.Join(root, "charlie/charlie.go")
	l.BeforeCommit(
		func() {
			assert.FatalOnError(
				t,
				os.WriteFile(charlie, []byte("package charlie\n\n"), 0644),
			)
		},
	)
	refused, e := l.Commit(nil)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{charlie}, refused)
	assert.String(
		t,
		"package alfa\n",
		testutil.ReadFile(t, filepath.Join(root, "alfa/alfa.go")),
	)
}

func TestLayerCommitPassesWriteOutsideReach(t *testing.T) {
	root := layerModule(t)
	l := layer.New(nil, root)
	l.Write(filepath.Join(root, "alfa/alfa.go"), []byte("package alfa\n\n"))
	l.BeforeCommit(
		func() {
			assert.FatalOnError(
				t,
				os.WriteFile(
					filepath.Join(root, "charlie/charlie.go"),
					[]byte("package charlie\n\n"),
					0644,
				),
			)
		},
	)
	refused, e := l.Commit([]*module_graph.Graph{module_graph.Build(root)})
	assert.FatalOnError(t, e)
	assert.Count(t, 0, refused)
	assert.String(
		t,
		"package alfa\n\n",
		testutil.ReadFile(t, filepath.Join(root, "alfa/alfa.go")),
	)
}

func TestLayerCommitRefusesWriteInImporter(t *testing.T) {
	root := layerModule(t)
	l := layer.New(nil, root)
	l.Write(filepath.Join(root, "alfa/alfa.go"), []byte("package alfa\n\n"))
	delta := filepath.Join(root, "delta/delta.go")
	l.BeforeCommit(
		func() {
			assert.FatalOnError(
				t,
				os.WriteFile(delta, []byte("package delta\n"), 0644),
			)
		},
	)
	refused, e := l.Commit([]*module_graph.Graph{module_graph.Build(root)})
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{delta}, refused)
}

func TestLayerCommitRefusesStaleDocument(t *testing.T) {
	root := layerModule(t)
	l := layer.New(nil, root)
	echo := filepath.Join(root, "doc/echo.md")
	l.Rewrite(echo, []byte("echo\n"), []byte("echo golf\n"))
	l.BeforeCommit(
		func() {
			assert.FatalOnError(
				t,
				os.WriteFile(echo, []byte("echo hotel\n"), 0644),
			)
		},
	)
	refused, e := l.Commit(nil)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{echo}, refused)
	assert.String(t, "echo hotel\n", testutil.ReadFile(t, echo))
}
