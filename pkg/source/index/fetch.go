package index

import "path"

func (w *Workspace) fetch(
	kind string,
	slot string,
	key string,
	fresh any,
) (any, bool) {
	k := path.Join(kind, slot)

	if e := w.memo[k]; e != nil && e.key == key {
		return e.value, true
	}

	if !w.store.Read(kind, key, fresh) {
		return nil, false
	}

	w.memo[k] = &memoEntry{key: key, value: fresh}

	return fresh, true
}
