package git

func IndexFiles(
	path string,
	keep func(string) bool,
) (map[string][]byte, error) {
	r, e := Open(path)

	if e != nil {
		return nil, e
	}

	scope, f := worktreeScope(r, path)

	if f != nil {
		return nil, f
	}

	x, g := r.Storer.Index()

	if g != nil {
		return nil, g
	}

	result := map[string][]byte{}

	for _, entry := range x.Entries {
		name, inside := scoped(entry.Name, scope)

		if !inside || !keep(name) {
			continue
		}

		content, h := blobContent(r, entry.Hash)

		if h != nil {
			return nil, h
		}

		result[name] = content
	}

	return result, nil
}
