package snapshot

import "github.com/funtimecoding/soil/pkg/source/types/snapshot_stamp"

type Snapshot struct {
	roots map[string]map[string]snapshot_stamp.Stamp
}
