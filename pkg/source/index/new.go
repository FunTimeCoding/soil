package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/generation"
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"github.com/funtimecoding/soil/pkg/source/index/store"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/types/file_sum"
	"github.com/funtimecoding/soil/pkg/source/types/memo_entry"
	"path/filepath"
	"time"
)

func New(
	directory string,
	root string,
	kinds ...*kind.Kind,
) *Workspace {
	now := time.Now()
	current := generation.Open(directory, generationName(), now, ownedEntry)
	build := extractorFingerprint()
	generation.Open(
		filepath.Join(current, constant.IndexBuildDirectory),
		buildName(build),
		now,
		ownedEntry,
	)

	return &Workspace{
		root:  root,
		graph: module_graph.Build(root),
		store: store.New(current),
		kinds: kinds,
		sums:  make(map[string]*file_sum.Sum),
		memo:  make(map[string]*memo_entry.Entry),
		build: build,
	}
}
