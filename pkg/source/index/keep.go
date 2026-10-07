package index

import (
	"github.com/funtimecoding/soil/pkg/source/types/memo_entry"
	"path"
)

func (w *Workspace) keep(
	kind string,
	slot string,
	key string,
	v any,
) {
	w.store.Write(kind, key, v)
	w.memo[path.Join(kind, slot)] = memo_entry.New(key, v)
}
