package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/types/top_entry"
	"maps"
)

func (t *Tree) liveStates(
	entries []*top_entry.Entry,
) (map[string]bool, map[string]bool) {
	reached, possible, missing := t.walk(entries)
	result := maps.Clone(reached)
	maps.Copy(result, possible)

	return result, missing
}
