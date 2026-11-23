package state_tree

func (t *Tree) Served() []string {
	live, _ := t.liveStates(t.entriesWhere(isLive))
	result := map[string]bool{}

	for s := range live {
		for _, p := range t.served[s] {
			result[p] = true
		}
	}

	return sortedKeys(result)
}
