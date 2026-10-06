package workspace

import "sort"

func (w *Workspace) Changes() []string {
	result := append([]string(nil), w.order...)
	sort.Strings(result)

	return result
}
