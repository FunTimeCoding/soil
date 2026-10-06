package index

import "path"

func (w *Workspace) holds(
	kind string,
	slot string,
	key string,
) bool {
	if e := w.memo[path.Join(kind, slot)]; e != nil && e.key == key {
		return true
	}

	return w.store.Has(kind, key)
}
