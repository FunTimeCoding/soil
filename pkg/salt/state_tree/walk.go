package state_tree

import "github.com/funtimecoding/soil/pkg/salt/types/top_entry"

func (t *Tree) walk(
	entries []*top_entry.Entry,
) (map[string]bool, map[string]bool, map[string]bool) {
	static, patterns := roots(entries)
	reached, missing := t.reachedStates(static)

	return reached, t.possibleStates(reached, patterns), missing
}
