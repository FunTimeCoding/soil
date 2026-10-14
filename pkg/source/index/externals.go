package index

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"slices"
)

func Externals[T any](
	w *Workspace,
	k *kind.Kind,
) []T {
	if k.External == nil && k.ExternalSyntax == nil {
		return nil
	}

	w.lock.Lock()
	defer w.lock.Unlock()
	values, missing := w.walkExternals(k)
	slices.Sort(missing)

	for len(missing) > 0 {
		w.loadExternals(missing, k.ExternalSyntax != nil)
		var next []string
		values, next = w.walkExternals(k)
		slices.Sort(next)

		if slices.Equal(next, missing) {
			break
		}

		missing = next
	}

	result := make([]T, 0, len(values))

	for _, v := range values {
		result = append(result, v.(T))
	}

	return result
}
