package reach

func ChangedEdits(
	root string,
	edits *Edits,
) ([]string, error) {
	changed, e := Changed(root)

	if e != nil {
		return nil, e
	}

	var result []string

	for _, p := range edits.Files {
		if changed[p] {
			result = append(result, p)
		}
	}

	return result, nil
}
