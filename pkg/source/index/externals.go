package index

import "github.com/funtimecoding/soil/pkg/source/index/kind"

func Externals[T any](
	w *Workspace,
	k *kind.Kind,
) []T {
	if k.External == nil {
		return nil
	}

	w.lock.Lock()
	defer w.lock.Unlock()
	values, missing := w.walkExternals(k)

	if len(missing) > 0 {
		w.loadExternals(missing)
		values, _ = w.walkExternals(k)
	}

	result := make([]T, 0, len(values))

	for _, v := range values {
		result = append(result, v.(T))
	}

	return result
}
