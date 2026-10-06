package snapshot

func Take(roots ...string) *Snapshot {
	result := &Snapshot{roots: make(map[string]map[string]stamp, len(roots))}

	for _, r := range roots {
		result.roots[r] = walk(r)
	}

	return result
}
