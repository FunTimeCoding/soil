package index

import "path"

func (w *Workspace) keep(
	kind string,
	slot string,
	key string,
	v any,
) {
	w.store.Write(kind, key, v)
	w.memo[path.Join(kind, slot)] = &memoEntry{key: key, value: v}
}
