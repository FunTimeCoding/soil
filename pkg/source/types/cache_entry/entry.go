package cache_entry

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
)

type Entry struct {
	Trees     []string
	Snapshot  *snapshot.Snapshot
	Workspace *index.Workspace
}
