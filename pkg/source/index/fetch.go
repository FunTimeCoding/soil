package index

import (
	"github.com/funtimecoding/soil/pkg/source/types/memo_entry"
	"path"
)

func (w *Workspace) fetch(
	kind string,
	slot string,
	key string,
	fresh any,
) (any, bool) {
	k := path.Join(kind, slot)

	if e := w.memo[k]; e != nil && e.Key == key {
		return e.Value, true
	}

	if !w.store.Read(kind, key, fresh) {
		return nil, false
	}

	w.memo[k] = memo_entry.New(key, fresh)

	return fresh, true
}
