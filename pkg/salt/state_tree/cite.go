package state_tree

func (t *Tree) Cite(paths []string) {
	for _, p := range paths {
		static, pattern := referenceOf(p)

		if pattern != nil {
			t.citedPattern = append(t.citedPattern, pattern)

			continue
		}

		t.cited = append(t.cited, static)
	}
}
