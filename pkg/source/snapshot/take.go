package snapshot

import "github.com/funtimecoding/soil/pkg/source/types/snapshot_stamp"

func Take(roots ...string) *Snapshot {
	result := &Snapshot{
		roots: make(map[string]map[string]snapshot_stamp.Stamp, len(roots)),
	}

	for _, r := range roots {
		result.roots[r] = walk(r)
	}

	return result
}
