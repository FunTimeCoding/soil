package reach

import "sort"

func ChangedEdits(
	root string,
	edits *Edits,
) ([]string, error) {
	changed, e := Changed(root)

	if e != nil {
		return nil, e
	}

	kept := map[string]bool{}

	for _, p := range edits.Files {
		if changed[p] {
			kept[p] = true
		}
	}

	for p := range changed {
		if matchesAny(edits.Patterns, p) {
			kept[p] = true
		}
	}

	var result []string

	for p := range kept {
		result = append(result, p)
	}

	sort.Strings(result)

	return result, nil
}
