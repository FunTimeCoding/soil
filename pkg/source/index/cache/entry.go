package cache

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/snapshot"
)

type entry struct {
	trees     []string
	snapshot  *snapshot.Snapshot
	workspace *index.Workspace
}
