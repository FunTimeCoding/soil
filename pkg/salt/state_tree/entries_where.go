package state_tree

import "github.com/funtimecoding/soil/pkg/salt/types/top_entry"

func (t *Tree) entriesWhere(
	keep func(e *top_entry.Entry) bool,
) []*top_entry.Entry {
	var result []*top_entry.Entry

	for _, e := range t.entries {
		if keep(e) {
			result = append(result, e)
		}
	}

	return result
}
