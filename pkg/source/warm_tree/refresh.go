package warm_tree

import "slices"

func (t *Tree) Refresh() []string {
	entries, children := walk(t.root)
	var result []string

	for p, e := range entries {
		if !e.Same(t.entries[p]) {
			result = append(result, p)
		}
	}

	for p := range t.entries {
		if _, kept := entries[p]; !kept {
			result = append(result, p)
		}
	}

	for _, p := range result {
		delete(t.content, p)
	}

	t.entries = entries
	t.children = children
	slices.Sort(result)

	return result
}
