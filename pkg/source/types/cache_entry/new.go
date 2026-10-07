package cache_entry

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
)

func New(
	trees []string,
	snapshot *snapshot.Snapshot,
	workspace *index.Workspace,
) *Entry {
	return &Entry{Trees: trees, Snapshot: snapshot, Workspace: workspace}
}
